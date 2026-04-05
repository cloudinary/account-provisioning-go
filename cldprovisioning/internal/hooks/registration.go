package hooks

func initHooks(h *Hooks) {
	cfg := NewAccountConfig()
	accountHook := NewCloudinaryAccountHook(cfg)

	h.registerSDKInitHook(accountHook)
	h.registerBeforeRequestHook(accountHook)
	h.registerSDKInitHook(&UserAgentHook{})
}
