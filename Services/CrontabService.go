package Services

import (
	"errors"
	"github.com/robfig/cron/v3"
	"sync"
)

type Crontab struct {
	inner *cron.Cron
	ids   map[string]cron.EntryID
	mutex sync.Mutex
}

func NewCrontab() *Crontab {
	return &Crontab{
		inner: cron.New(),
		ids:   make(map[string]cron.EntryID),
	}
}

func (c *Crontab) IDs() []string {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	validIDs := make([]string, 0, len(c.ids))
	invalidIDs := make([]string, 0)
	for sid, eid := range c.ids {
		if e := c.inner.Entry(eid); e.ID != eid {
			invalidIDs = append(invalidIDs, sid)
			continue
		}
		validIDs = append(validIDs, sid)
	}
	for _, id := range validIDs {
		delete(c.ids, id)
	}
	return validIDs
}

func (c *Crontab) Start() {
	c.inner.Start()
}

func (c *Crontab) Stop() {
	c.inner.Stop()
}

func (c *Crontab) DelById(id string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	eid, ok := c.ids[id]
	if !ok {
		return
	}
	c.inner.Remove(eid)
	delete(c.ids, id)
}

func (c *Crontab) AddByID(id string, spec string, cmd cron.Job) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.ids[id]; ok {
		return errors.New("crontab id exists")
	}
	eid, err := c.inner.AddJob(spec, cmd)
	if err != nil {
		return err
	}
	c.ids[id] = eid
	return nil
}

func (c *Crontab) AddByFunc(id string, spec string, f func()) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.ids[id]; ok {
		return errors.New("crontab id exists")
	}

	eid, err := c.inner.AddFunc(spec, f)
	if err != nil {
		return err
	}
	c.ids[id] = eid
	return nil
}

// @Summer 判断任务十分存在
func (c *Crontab) IsExists(id string) bool {
	_, ok := c.ids[id]
	return ok
}
