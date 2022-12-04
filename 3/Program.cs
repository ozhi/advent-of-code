// See https://aka.ms/new-console-template for more information

const int MAX_PRIORITY = 1 + 26 + 26; 

int priority(char c) {
    if ('a' <= c && c <= 'z') {
        return 1 + c - 'a';
    }
    return 1 + 26 + c - 'A';
}

void part1() {
    int totalPriority = 0;
    String s;
    while((s = Console.ReadLine()) != null) {
        bool[] firstHalf = new bool[MAX_PRIORITY];
        bool[] secondHalf = new bool[MAX_PRIORITY];
        for(int i = 0; i < s.Length; i++) {
            if (i < s.Length/2) {
                firstHalf[priority(s[i])] = true;
            } else {
                secondHalf[priority(s[i])] = true;
            }
        }

        for (int i = 1; i <= MAX_PRIORITY; i++) {
            if (firstHalf[i] && secondHalf[i]) {
                totalPriority += i;
                break; // Only one shared item per rucksack.
            }
        }

        Console.WriteLine("totalPriority: {0}", totalPriority);
    }
}

void markChars(string s, bool[] map) {
    for(int i = 0; i < s.Length; i++) {
        map[priority(s[i])] = true;
    }
}

void part2() {
    int totalPriority = 0;
    String s;
    while(true) {
        bool[] first = new bool[MAX_PRIORITY];
        bool[] second = new bool[MAX_PRIORITY];
        bool[] third = new bool[MAX_PRIORITY];
        
        s = Console.ReadLine();
        markChars(s, first);

        s = Console.ReadLine();
        markChars(s, second);

        s = Console.ReadLine();
        markChars(s, third);

        
        for(int i = 1; i <= MAX_PRIORITY; i++) {
            if (first[i] && second[i] && third[i]) {
                Console.WriteLine("shared item: {0}", i);
                totalPriority += i;
                break;
            }
        }

        Console.WriteLine("totalPriority: {0}", totalPriority);
    }
}

part2();