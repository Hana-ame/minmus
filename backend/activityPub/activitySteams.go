package activityPub

type activityStream map[string]any

func (a *activityStream) Get(key string) (any, bool) {
	return a, false
}
