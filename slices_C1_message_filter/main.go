package main

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line

func filterMessages(messages []Message, filterType string) []Message {
	var filteredMessages []Message // slice dari message // untuk menyimpan pesan yang sudah difilter berdasarkan tipe yang diinginkan.

	for _, msg := range messages { // Loop melalui setiap pesan dalam slice input
		if msg.Type() == filterType { // Periksa tipe pesan untuk setiap elemen dalam loop
			filteredMessages = append(filteredMessages, msg) // Tambahkan pesan ke slice hasil jika tipe pesan cocok dengan tipe yang diinginkan
		}
	}

	return filteredMessages
}
