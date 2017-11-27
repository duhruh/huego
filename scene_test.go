package huego

import (
	"testing"
	"os"
)

func TestGetScenes(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	scenes, err := hue.GetScenes()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d scenes", len(scenes))
	for _, scene := range scenes {
		t.Logf("Scene id=%d name=%s", scene.Id, scene.Name)
	}
}

func TestGetScene(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	scenes, err := hue.GetScenes()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d scenes", len(scenes))
	for _, scene := range scenes {
		t.Logf("Getting scene %d, skipping the rest", scene.Id)
		s, err := hue.GetScene(scene.Id)
		if err != nil {
			t.Error(err)
		}
		t.Logf("Got scene name=%s", s.Name)
		break
	}
}

func TestCreateScene(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	newScene := &Scene{Name: "TestScene"}
	response, err := hue.CreateScene(newScene)
	if err != nil {
		t.Error(err)
	}
	for _, r := range response {
		t.Logf("Response from put: Success=%v Error=%v", r.Success, r.Error)
	}
}


func TestUpdateScene(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	scenes, err := hue.GetScenes()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d scenes, setting the first one", len(scenes))
	for _, scene := range scenes {
		response, err := hue.UpdateScene(scene.Id, scene)
		if err != nil {
			t.Error(err)
		}
		for _, r := range response {
			t.Logf("Response from put: Success=%v Error=%v", r.Success, r.Error)
		}
		break
	}
}

// func TestDeleteScene(t *testing.T) {
// 	res, err := hue.DeleteScene(1)
// 	if err != nil {
// 		t.Log(err)
// 		t.Fail()
// 	} else {
// 		for _, r := range res {
// 			t.Log(r.Success, r.Error)
// 		}
// 	}
// }
