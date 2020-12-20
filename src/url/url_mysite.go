/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package url

const PREFIX_MYSITE = "/rest/mysite"

const(
	MYSITE_INDEX = ""
	MYSITE_GET_BLOG = "/get_blog"
	MYSITE_GET_POST = "/get_post"
	MYSITE_GET_PROBLEMS = "/get_problems"
	MYSITE_GET_THOUGHTS = "/get_thoughts"
	MYSITE_GET_MESSAGE = "/get_message"
	MYSITE_GET_MUSIC = "/get_music"
	MYSITE_GET_VIEWS = "/get_views"
)

func URL2MYSITE(raw string) string {
	return PREFIX_MYSITE + raw
}