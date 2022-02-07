package tik_lib

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
	setdata_common "github.com/kirigaikabuto/setdata-common"
)

type AmqpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewAmqpEndpoints(ch setdata_common.CommandHandler) AmqpEndpoints {
	return AmqpEndpoints{ch: ch}
}

func (a *AmqpEndpoints) CreateUser() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &CreateUserCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := a.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}

func (a *AmqpEndpoints) UpdateUser() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &UpdateUserCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := a.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}

func (a *AmqpEndpoints) GetUserById() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &GetUserByIdCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := a.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}

func (a *AmqpEndpoints) ListUser() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &ListUserCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := a.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}

func (a *AmqpEndpoints) DeleteUser() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &DeleteUserCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := a.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}

func (a *AmqpEndpoints) GetUserByPhoneNumber() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &GetUserByPhoneNumberCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := a.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}

func (a *AmqpEndpoints) CreateFile() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &CreateFileCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := a.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}

func (a *AmqpEndpoints) UpdateFile() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &UpdateFileCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := a.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}

func (a *AmqpEndpoints) GetFileById() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &GetFileByIdCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := a.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}

func (a *AmqpEndpoints) DeleteFile() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &DeleteFileCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := a.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}

func (a *AmqpEndpoints) ListFiles() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &ListFilesCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := a.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}





