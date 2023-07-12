package tg_message_constructor

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/Red-Sock/gitm8/internal/service/domain"
	"github.com/Red-Sock/gitm8/internal/transport/tg/assets"
)

type workflowConstructor struct {
}

func (m *MessageConstructor) extractWorkflowRun(payload domain.Payload) (string, []tgbotapi.MessageEntity, error) {
	// TODO GITM-28: now - sending message on every action / After - send 1 message and then change it

	switch payload.GetWorkflow().Status {
	case domain.WorkflowStatusInProgress:
		return m.workflowConstructor.extractWorkflowInProgress(payload)

	case domain.WorkflowStatusCompleted:
		return m.workflowConstructor.extractWorkflowCompleted(payload)

	default:
		return "", nil, domain.ErrInvalidActionForEvent
	}

}
func (w *workflowConstructor) extractWorkflowInProgress(payload domain.Payload) (string, []tgbotapi.MessageEntity, error) {
	constr := constructor{}

	constr.Write(assets.Loading)
	constr.Write("Pipeline ")
	{
		wf := payload.GetWorkflow()
		constr.WriteWithLink(wf.Name, wf.Link)
	}
	constr.Write(" has started ")

	constr.Writeln(" Project: ")
	{
		proj := payload.GetProject()
		constr.WriteWithLink(proj.Name, proj.Link)
	}

	return constr.String(), constr.format, nil
}

func (w *workflowConstructor) extractWorkflowCompleted(payload domain.Payload) (string, []tgbotapi.MessageEntity, error) {
	constr := constructor{}

	switch payload.GetWorkflow().Result {
	case domain.WorkflowResultSuccess:
		constr.Write(assets.Completed)
	case domain.WorkflowResultFailure:
		constr.Write(assets.Failed)
	}

	constr.Write("Pipeline ")
	{
		wf := payload.GetWorkflow()
		constr.WriteWithLink(wf.Name, wf.Link)
	}
	constr.Write(" has finished ")

	constr.Writeln("Project: ")
	{
		proj := payload.GetProject()
		constr.WriteWithLink(proj.Name, proj.Link)
	}

	return constr.String(), constr.format, nil
}
