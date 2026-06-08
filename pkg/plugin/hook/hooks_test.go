package hook

import "testing"

func TestSceneOUpdatePost_IsValidAndString(t *testing.T) {
	if SceneOUpdatePost != "Scene.OUpdate.Post" {
		t.Fatalf("got %q, want Scene.OUpdate.Post", SceneOUpdatePost)
	}
	if !SceneOUpdatePost.IsValid() {
		t.Fatal("SceneOUpdatePost should be a valid trigger")
	}
	found := false
	for _, e := range AllHookTriggerEnum {
		if e == SceneOUpdatePost {
			found = true
		}
	}
	if !found {
		t.Fatal("SceneOUpdatePost missing from AllHookTriggerEnum")
	}
}
