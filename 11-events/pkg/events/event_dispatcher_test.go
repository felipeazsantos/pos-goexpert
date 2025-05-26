package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}
func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface) {

}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3		TestEventHandler
	eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatcher = NewEventDispatcher()
	suite.handler = TestEventHandler{ID: 1}
	suite.handler2 = TestEventHandler{ID: 2}
	suite.handler3 = TestEventHandler{ID: 3}
	suite.event = TestEvent{Name: "test.event", Payload: "test.payload"}
	suite.event2 = TestEvent{Name: "test.event2", Payload: "test.payload2"}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)	
	assert.NoError(suite.T(), err)
    suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	assert.NoError(suite.T(), err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	suite.Equal(suite.eventDispatcher.handlers[suite.event.GetName()][0], &suite.handler)
	suite.Equal(suite.eventDispatcher.handlers[suite.event.GetName()][1], &suite.handler2)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_AlreadyRegistered() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.NoError(err)
    suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Error(err)
	suite.Equal("handler already registered", err.Error())
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
	suite.Equal(suite.eventDispatcher.handlers[suite.event.GetName()][0], &suite.handler)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.NoError(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.NoError(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))


	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.NoError(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	suite.eventDispatcher.Clear()
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.NoError(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.NoError(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	suite.True(suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler))
	suite.True(suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler2))
	suite.False(suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler3))
}


func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
