bool containsFully(int from1, int to1, int from2, int to2) {
    return (from2 - from1) * (to1 - to2) >= 0;
    //     (from2 <= from1     && to1 <= to2)      || (from1 <= from2     && to2 <= to1)
    // <=> (from2 - from1 <= 0 && to1 - to2 <= 0 ) || (from1 - from2 <= 0 && to2 - to1 <= 0)
    // <=> (a<=0 && b<=0) || (a>=0 && b>=0)
    // <=> a*b >= 0
}

bool overlapAtAll(int from1, int to1, int from2, int to2) {
    return
        containsFully(from1, to1, from2, from2) ||
        containsFully(from1, from1, from2, to2);
}

void solve() {
    char[] delimiters = {',', '-'};
    int count1 = 0;
    int count2 = 0;

    foreach (string line in System.IO.File.ReadLines(@"./my-input.txt")) {
        string[] parts = line.Split(delimiters);

        int from1 = Int32.Parse(parts[0]);
        int to1 = Int32.Parse(parts[1]);
        int from2 = Int32.Parse(parts[2]);
        int to2 = Int32.Parse(parts[3]);
        
        if (containsFully(from1, to1, from2, to2)) {
            count1++;
        }
        if (overlapAtAll(from1, to1, from2, to2)) {
            count2++;
        }
    }
    Console.WriteLine("Part1: {0}\nPart2: {1}", count1, count2);
}

solve();
