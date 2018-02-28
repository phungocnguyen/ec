package models

import ()

type Control struct {
	Id          int
	ReqId       int
	CisId       string
	Category    string
	Requirement string
	Discussion  string
	CheckText   string
	FixText     string
	RowDesc     string
	BaselineId  int
}

func (c *Control) SetId(id int) {
	c.Id = id
}

func (c *Control) SetReqId(reqId int) {
	c.ReqId = reqId
}

func (c *Control) SetBaselineId(baselineId int) {
	c.BaselineId = baselineId
}

func (c *Control) SetCisId(cisId string) {
	c.CisId = cisId
}

func (c *Control) SetCategory(category string) {
	c.Category = category
}

func (c *Control) SetRequirement(requirement string) {
	c.Requirement = requirement
}

func (c *Control) SetDiscussion(discussion string) {
	c.Discussion = discussion
}

func (c *Control) SetCheckText(checkText string) {
	c.CheckText = checkText
}

func (c *Control) SetFixText(fixText string) {
	c.FixText = fixText
}

func (c *Control) SetRowDesc(rowDesc string) {
	c.RowDesc = rowDesc
}
