package main


import(
"fmt"
"strings"
"path/filepath"
)

/*

- maps are unordered lists
- maps key-value pairs
- maps are dynamically resizable
- maps are references
- all keys in the map have to be unique
- maps are reference types. Passed to functions as reference
Thus, changes made by functions, are visible by caller
- maps are unsafe for concurrency

syntax

map[<keyType>]<valueType>


specify the size of the maps

make(map[<keyType>]<valueType>, size)

* specifying a size can improve performance


*/

func main(){

	// displayMap1();

	//displayMap2();

	// modifyMap(); 

	setupInitLayer();

}


func displayMap1() {

	leagueTitles := make(map[string]int)

	leagueTitles["Sunderland"] = 6
	leagueTitles["Newcastle"] = 4

	recentHead2Head := map[string]int {
		"Sunderland": 5,
		"Newcastle": 0,
	}

	fmt.Printf("\nLeague titles %v\nRecent head 2 head %v\n",
		leagueTitles, recentHead2Head)

}


func displayMap2() {

	testMap := map[string]int{"A":1,"B":2,"C":3,"D":4,"E":5}

	for key, value := range testMap {
		fmt.Printf("\nKey:%v Value:%v", key, value)
	}


}


func modifyMap() {

    testMap := map[string]int{"A":1,"B":2,"C":3,"D":4,"E":5}

    testMap["A"] = 100
    testMap["F"] = 1900
	
	fmt.Println(testMap)

	//delete an item
	delete(testMap, "F")

	fmt.Println(testMap)
			
}



// setupInitLayer populates a directory with mountpoints suitable
// for bind-mounting things into the container.
//
// This extra layer is used by all containers as the top-most ro layer. It protects
// the container from unwanted side-effects on the rw layer.
func setupInitLayer() error {
	for pth, typ := range map[string]string{
		"/dev/pts":         "dir",
		"/dev/shm":         "dir",
		"/proc":            "dir",
		"/sys":             "dir",
		"/.dockerenv":      "file",
		"/etc/resolv.conf": "file",
		"/etc/hosts":       "file",
		"/etc/hostname":    "file",
		"/dev/console":     "file",
		"/etc/mtab":        "/proc/mounts",
	} {
		fmt.Println("Type:", typ)
		parts := strings.Split(pth, "/")
		prev := "/"
		for _, p := range parts[1:] {
			prev = filepath.Join(prev, p)

			fmt.Println("Prev:", prev)
			
		}

/*
		if _, err := os.Stat(filepath.Join(initLayer, pth)); err != nil {
			if os.IsNotExist(err) {
				if err := idtools.MkdirAllNewAs(filepath.Join(initLayer, filepath.Dir(pth)), 0755, rootUID, rootGID); err != nil {
					return err
				}
				switch typ {
				case "dir":
					if err := idtools.MkdirAllNewAs(filepath.Join(initLayer, pth), 0755, rootUID, rootGID); err != nil {
						return err
					}
				case "file":
					f, err := os.OpenFile(filepath.Join(initLayer, pth), os.O_CREATE, 0755)
					if err != nil {
						return err
					}
					f.Chown(rootUID, rootGID)
					f.Close()
				default:
					if err := os.Symlink(typ, filepath.Join(initLayer, pth)); err != nil {
						return err
					}
				}
			} else {
				return err
			}
		}
		*/
	}


	// Layer is ready to use, if it wasn't before.
	return nil
}