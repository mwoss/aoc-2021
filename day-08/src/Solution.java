import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Arrays;
import java.util.List;

public class Solution {

    public long countUniquelyDistinguishableSignal(String rawDisplayData) throws IOException {
        return Arrays.stream(rawDisplayData.trim().split("\n"))
                .map(line -> line.split(" \\| ")[1].split(" "))
                .flatMap(Arrays::stream)
                .filter(this::isUniquelyDistinguishableSignal)
                .count();
    }

    private boolean isUniquelyDistinguishableSignal(String signal) {
        // digits 1, 4, 7, and 8 each use a unique number of segments
        int signalSize = signal.length();
        return signalSize == 2 || signalSize == 3 || signalSize == 4 || signalSize == 7;
    }


    public static void main(String[] args) throws IOException {
        String rawDisplayData = Files.readString(Path.of("src/input.txt"));

        Solution solution = new Solution();
        System.out.println(solution.countUniquelyDistinguishableSignal(rawDisplayData));
    }
}
