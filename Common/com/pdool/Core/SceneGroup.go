package Core

import "com/pdool/DataStruct"

type SceneGroup struct {
	groupId int			//	分线id
	sceneId int 		//	场景id
	playerList 		DataStruct.LinkedList
	otherList		DataStruct.LinkedList
}