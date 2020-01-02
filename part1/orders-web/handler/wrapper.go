package handler

import (
	"context"
	"net/http"

	auth "auth/proto/auth"
	"basic/common"
	"github.com/micro/go-micro/util/log"
	"plugins/session"
)

// AuthWrapper 认证wrapper
func AuthWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ck, _ := r.Cookie(common.RememberMeCookieName)
		// token不存在，则状态异常，无权限
		if ck == nil {
			http.Error(w, "非法请求", 400)
			return
		}

		sess := session.GetSession(w, r)
		if sess.ID == "" {
			http.Error(w, "非法请求", 400)
			return
		}

		// 检测是否通过验证
		if sess.Values["valid"] != nil {
			h.ServeHTTP(w, r)
			return
		}

		userId := sess.Values["userId"].(int64)
		if userId == 0 {
			log.Logf("[AuthWrapper]，session不合法，无用户id")
			http.Error(w, "非法请求", 400)
			return
		}

		rsp, err := authClient.GetCachedAccessToken(context.TODO(), &auth.Request{
			UserId: userId,
		})

		if err != nil {
			log.Logf("[AuthWrapper]，err：%s", err)
			http.Error(w, "非法请求", 400)
			return
		}

		// token不一致
		if rsp.Token != ck.Value {
			log.Logf("[AuthWrapper]，token不一致")
			http.Error(w, "非法请求", 400)
			return
		}

		h.ServeHTTP(w, r)
	})
}
