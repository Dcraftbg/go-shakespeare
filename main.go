package main
import "os"
import "log"
import "fmt"
import "strings"
func main() {
    res, err := os.ReadFile("./shakespeare.txt");
    if err != nil {
        log.Fatal("ERROR: ", err);
    }
    res_str := string(res);
    res_iter := strings.FieldsSeq(res_str);
    biggest_word_len := 0;
    freq := make(map[string]int);
    for s := range res_iter {
        if biggest_word_len <= len(s) {
            biggest_word_len = len(s);
        }
        freq[s] += 1;
    }
    most_commonly_used_word := "";
    most_commonly_used_freq := 0;
    for word, freq := range freq {
        if most_commonly_used_freq < freq { 
            most_commonly_used_freq = freq;
            most_commonly_used_word = word;
        }
        // fmt.Println(k, s);
    }
    fmt.Println("Most commonly encountered word:", most_commonly_used_word, "with a frequency of", most_commonly_used_freq);
}
