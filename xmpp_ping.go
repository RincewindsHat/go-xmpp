package xmpp

import (
	"fmt"
)

func (c *Client) PingC2S(jid, server string) error {
	if jid == "" {
		jid = c.jid
	}
	if server == "" {
		server = c.domain
	}
	_, err := fmt.Fprintf(c.stanzaWriter, "<iq from='%s' to='%s' id='%s' type='get'>"+
		"<ping xmlns='urn:xmpp:ping'/>"+
		"</iq>\n",
		xmlEscape(jid), xmlEscape(server), getUUIDv4())
	return err
}

func (c *Client) PingS2S(fromServer, toServer string) error {
	_, err := fmt.Fprintf(c.stanzaWriter, "<iq from='%s' to='%s' id='%s' type='get'>"+
		"<ping xmlns='urn:xmpp:ping'/>"+
		"</iq>\n",
		xmlEscape(fromServer), xmlEscape(toServer), getUUIDv4())
	return err
}

func (c *Client) SendResultPing(id, toServer string) error {
	_, err := fmt.Fprintf(c.stanzaWriter, "<iq type='result' to='%s' id='%s'/>\n",
		xmlEscape(toServer), xmlEscape(id))
	return err
}
