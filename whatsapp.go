package componentes

type Contatos struct {
	Celular     string `json:"contato"`
	Cliente     string `json:"cliente"`
	Notificacao string `json:"notificacao"`
}

type WhatsappMensagem struct {
	ExpiraLista       int        `json:"expiraLista"`
	CancelarPendentes bool       `json:"cancelarPendentes"`
	Contato           []Contatos `json:"contatos"`
}

// Envio Mensagem de Whatsapp usando o Broken da Escallo
func WhatsAppEnvio(webService string, token string, paciente string, celular string, mensagem string) error {
	var whatsapp WhatsappMensagem
	contato := []Contatos{Contatos{Celular: celular, Cliente: paciente, Notificacao: mensagem}}
	whatsapp.CancelarPendentes = false
	whatsapp.ExpiraLista = 60
	whatsapp.Contato = contato
	//	endpoint := conect.Webservice + "escallo/api/v1/campanha/texto/18/lista"
	endpoint := webService
	identificador := "whatsapp"
	_, err := Wenvio(&whatsapp, endpoint, identificador, token)
	return err
}
