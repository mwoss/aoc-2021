import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.*;
import java.util.function.Predicate;
import java.util.stream.Collectors;

public class Solution {

    public long countUniquelyDistinguishableSignal(String rawDisplayData) {
        return Arrays.stream(rawDisplayData.trim().split("\n"))
                .map(line -> line.split(" \\| ")[1].split(" "))
                .flatMap(Arrays::stream)
                .filter(this::isUniquelyDistinguishableSignal)
                .count();
    }

    public long sumDecodedOutputSignals(String rawDisplayData) {
        var x = Arrays.stream(rawDisplayData.trim().split("\n"))
                .map(line -> line.split(" \\| ")).toList();

        var overallSum = 0L;
        for (String[] line : x) {
            var left = line[0];
            var right = line[1];

            var numbers = left.split(" ");
            var distinguishableNumbers = Arrays.stream(numbers).filter(this::isUniquelyDistinguishableSignal).sorted(Comparator.comparingInt(String::length)).toList(); // 1,7,4,8
            var indistinguishableNumbers = Arrays.stream(numbers).filter(Predicate.not(this::isUniquelyDistinguishableSignal)).toList();

            var knownNumbers = Map.of(
                    1, toSet(distinguishableNumbers.get(0)),
                    7, toSet(distinguishableNumbers.get(1)),
                    4, toSet(distinguishableNumbers.get(2)),
                    8, toSet(distinguishableNumbers.get(3))
            );

            var signalToNumber = new HashMap<Set<Character>, Integer>(Map.of(
                    toSet(distinguishableNumbers.get(0)), 1,
                    toSet(distinguishableNumbers.get(1)), 7,
                    toSet(distinguishableNumbers.get(2)), 4,
                    toSet(distinguishableNumbers.get(3)), 8
            ));

            for (String indistinguishableNumber : indistinguishableNumbers) {
                signalToNumber.put(toSet(indistinguishableNumber), decodeSignal(indistinguishableNumber, knownNumbers));
            }

            var output = right.split(" ");
            var multiplier = 1;
            var number = 0;
            for (int i = output.length - 1; i >= 0; i--) {
                number += (long) signalToNumber.get(toSet(output[i])) * multiplier;
                multiplier *= 10;
            }
            System.out.println(number);
            overallSum += number;
        }

        return overallSum;
    }

    private int decodeSignal(String signal, Map<Integer, Set<Character>> baseMapping) {
        var signalLen = signal.length();

        var signalCharacters = signal.chars().mapToObj(e -> (char) e).collect(Collectors.toSet());

        if (signalLen == 5) {
            var one = baseMapping.get(1);
            var eightMinusFour = new HashSet<>(baseMapping.get(8));
            eightMinusFour.removeAll(baseMapping.get(4));
            if (signalCharacters.stream().filter(one::contains).collect(Collectors.toSet()).equals(one)) {
                return 3;
            } else if (signalCharacters.stream().filter(eightMinusFour::contains).collect(Collectors.toSet()).equals(eightMinusFour)) {
                return 2;
            } else {
                return 5;
            }
        } else {
            var eightMinusSeven = new HashSet<>(baseMapping.get(8));
            eightMinusSeven.removeAll(baseMapping.get(7));
            var fourMinusOne = new HashSet<>(baseMapping.get(4));
            fourMinusOne.removeAll(baseMapping.get(1));


            System.out.print(baseMapping.get(8));
            System.out.print(baseMapping.get(7));
            System.out.print(signalCharacters);
            System.out.print(eightMinusSeven);
            System.out.println(signalCharacters.stream().filter(eightMinusSeven::contains).collect(Collectors.toSet()));

            if (signalCharacters.stream().filter(eightMinusSeven::contains).collect(Collectors.toSet()).equals(eightMinusSeven)) {
                return 6;
            } else if (signalCharacters.stream().filter(fourMinusOne::contains).collect(Collectors.toSet()).equals(fourMinusOne)) {
                return 9;
            } else {
                return 0;
            }
        }
    }

    private Set<Character> toSet(String input) {
        return input.chars().mapToObj(e -> (char) e).collect(Collectors.toSet());
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
        System.out.println(solution.sumDecodedOutputSignals(rawDisplayData));
    }
}
