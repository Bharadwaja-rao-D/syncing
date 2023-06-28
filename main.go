package main

import (
	//"os"

	"github.com/bharadwaja-rao-d/syncing/client"
	//"github.com/bharadwaja-rao-d/syncing/diff"
)

func main() {

	//f1, _ := os.ReadFile("./test/file1.txt")
	//f2, _ := os.ReadFile("./test/file2.txt")

    client := client.NewClient();
    client.Watch("./test/file1.txt");

    /*
	edit_script := diff.Differ(string(f1), string(f2))
	fmt.Println(edit_script)

	dst := diff.DeDiffer(string(f1), edit_script)
	fmt.Println(dst)
    */
}
