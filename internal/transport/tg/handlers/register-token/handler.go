package register_token

// TODO when tokens come in to play
//const Command = "/register"
//
//type MainMenu struct {
//	interfaces.RegistrationService
//}
//
//func New(regSrv interfaces.RegistrationService) *MainMenu {
//	return &MainMenu{
//		RegistrationService: regSrv,
//	}
//}
//
//func (m *MainMenu) Handle(in *model.MessageIn, out tginterfaces.Chat) {
//	req, err := m.obtainGitSystem(in, out)
//	if err != nil {
//		logrus.Errorf("error obtaining git system for user_id: %d, in chat %d - %s", in.Contact.UserID, in.Chat.ID, err.Error())
//		return
//	}
//
//	err = m.obtainToken(in, out, &req)
//	if err != nil {
//		logrus.Errorf("error obtaining git token for user_id: %d, in chat %d - %s", in.Contact.UserID, in.Chat.ID, err.Error())
//		return
//	}
//
//	err = m.RegistrationService.CreateTicket(context.Background(), req)
//	if err != nil {
//		// TODO
//		out.SendMessage(response.NewMessage(err.Error()))
//		err = errors.Join(err, fmt.Errorf("error registrating git system for user_id: %d, in chat %d - %s", in.Contact.UserID, in.Chat.ID, err.Error()))
//		logrus.Error(err.Error())
//		return
//	}
//}
//
//func (m *MainMenu) obtainGitSystem(srcMsg *model.MessageIn, out tginterfaces.Chat) (req domain.RegisterRequest, err error) {
//	msg := &response.EditMessage{
//		ChatId:    srcMsg.Chat.ID,
//		MessageId: int64(srcMsg.MessageID),
//		Text:      "Input url or select one of the following git system:",
//	}
//	msg.Keys = &menu.InlineKeyboard{}
//
//	msg.Keys.AddButton("github", "github")
//
//	msg.MessageId = int64(srcMsg.MessageID)
//	out.SendMessage(msg)
//
//	in, err := out.GetInput(context.Background())
//	if err != nil {
//		return req, err
//	}
//
//	URL, err := url.Parse(in.Text)
//	if err == nil && URL.Host != "" {
//		req.RepositoryURl = URL.Host
//	} else {
//		err = req.RepoType.SetType(in.Text)
//		if err != nil {
//			return req, err
//		}
//	}
//
//	return req, nil
//}
//
//func (m *MainMenu) obtainToken(srcMsg *model.MessageIn, out tginterfaces.Chat, req *domain.RegisterRequest) error {
//	msg := &response.EditMessage{
//		Text: "Enter your git token",
//	}
//	const docsCommand = "token docs"
//	{
//		msg.ForceSetMessageId(int64(srcMsg.MessageID))
//		kb := &menu.InlineKeyboard{}
//
//		kb.AddButton("Where can I get token?", docsCommand)
//
//		msg.Keys = kb
//	}
//	out.SendMessage(msg)
//
//	for {
//		ctx, _ := context.WithTimeout(context.Background(), time.Minute*5)
//
//		resp, err := out.GetInput(ctx)
//		if err != nil {
//			return err
//		}
//
//		switch resp.Text {
//		case docsCommand:
//			helpMsg := req.RepoType.GetHelpMessage()
//			out.SendMessage(response.NewMessage("[generate token](" + helpMsg + ")"))
//			continue
//		default:
//			req.Token = resp.Text
//			return nil
//		}
//	}
//}
