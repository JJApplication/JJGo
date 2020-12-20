/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package url

const PREFIX_JJMAIL = "/rest/jjmail"

const(
	JJMAIL_INDEX = ""
	JJMAIL_STATUS = "/status"
	JJMAIL_SUB_BLOG = "/sub_blog"
	JJMAIL_UNSUB_BLOG = "/unsub_blog"
	JJMAIL_SUB_MGEK = "/sub_mgek"
	JJMAIL_UNSUB_MGEK = "/unsub_mgek"
	JJMAIL_SEND = "/send"
	JJMAIL_REPLY = "/reply"
)

func URL2JJMAIL(raw string) string {
	return PREFIX_JJMAIL + raw
}