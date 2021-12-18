import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.*;
import java.util.function.Predicate;
import java.util.stream.Collectors;

public class Solution {

    private record Entry(String[] inputSignal, String[] outputSignals) {
        static Entry parse(String line) {
            var signals = line.split(" \\| ");
            return new Entry(signals[0].split(" "), signals[1].split(" "));
        }
    }

    private static class CharacterSet {
        private static Set<Character> of(String input) {
            return input.chars().mapToObj(e -> (char) e).collect(Collectors.toUnmodifiableSet());
        }
    }

    public long countUniquelyDistinguishableSignal(String rawDisplayData) {
        return Arrays.stream(rawDisplayData.trim().split("\n"))
                .map(line -> line.split(" \\| ")[1].split(" "))
                .flatMap(Arrays::stream)
                .filter(this::isUniquelyDistinguishableSignal)
                .count();
    }

    public long sumDecodedOutputSignals(String rawDisplayData) {
        var entries = Arrays.stream(rawDisplayData.trim().split("\n")).map(Entry::parse).toList();

        var overallSum = 0L;
        for (Entry entry : entries) {
            var distinguishableSignals = Arrays.stream(entry.inputSignal)
                    .filter(this::isUniquelyDistinguishableSignal)
                    .sorted(Comparator.comparingInt(String::length))
                    .toList(); // guaranteed that contains only four signals (ordered): 1,7,4,8
            var indistinguishableSignals = Arrays.stream(entry.inputSignal)
                    .filter(Predicate.not(this::isUniquelyDistinguishableSignal))
                    .toList();


            var knownNumberToDecomposedSignal = Map.of(
                    1, CharacterSet.of(distinguishableSignals.get(0)),
                    7, CharacterSet.of(distinguishableSignals.get(1)),
                    4, CharacterSet.of(distinguishableSignals.get(2)),
                    8, CharacterSet.of(distinguishableSignals.get(3))
            );

            var decomposedSignalToNumber = new HashMap<>(Map.of(
                    CharacterSet.of(distinguishableSignals.get(0)), 1,
                    CharacterSet.of(distinguishableSignals.get(1)), 7,
                    CharacterSet.of(distinguishableSignals.get(2)), 4,
                    CharacterSet.of(distinguishableSignals.get(3)), 8
            ));

            for (String signal : indistinguishableSignals) {
                decomposedSignalToNumber.put(CharacterSet.of(signal), decodeSignal(signal, knownNumberToDecomposedSignal));
            }

            overallSum += calculateOutputValue(entry.outputSignals, decomposedSignalToNumber);
        }

        return overallSum;
    }

    private int decodeSignal(String signal, Map<Integer, Set<Character>> knownNumberToDecomposedSignal) {
        var decomposedSignal = CharacterSet.of(signal);

        if (signal.length() == 5) {
            var eightMinusFour = new HashSet<>(knownNumberToDecomposedSignal.get(8));
            eightMinusFour.removeAll(knownNumberToDecomposedSignal.get(4));

            if (decomposedSignal.containsAll(knownNumberToDecomposedSignal.get(1))) {
                return 3;
            }
            if (decomposedSignal.containsAll(eightMinusFour)) {
                return 2;
            }
            return 5;
        }

        var eightMinusSeven = new HashSet<>(knownNumberToDecomposedSignal.get(8));
        eightMinusSeven.removeAll(knownNumberToDecomposedSignal.get(7));

        var fourMinusOne = new HashSet<>(knownNumberToDecomposedSignal.get(4));
        fourMinusOne.removeAll(knownNumberToDecomposedSignal.get(1));

        if (decomposedSignal.containsAll(eightMinusSeven)) {
            return 6;
        }
        if (decomposedSignal.containsAll(fourMinusOne)) {
            return 9;
        }
        return 0;
    }

    private int calculateOutputValue(String[] outputSignals, HashMap<Set<Character>, Integer> signalToNumber) {
        var multiplier = 1;
        var number = 0;
        for (int i = outputSignals.length - 1; i >= 0; i--) {
            number += (long) signalToNumber.get(CharacterSet.of(outputSignals[i])) * multiplier;
            multiplier *= 10;
        }
        return number;
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
