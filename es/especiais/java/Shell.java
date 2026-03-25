import java.util.Scanner;
import java.util.List;
import java.util.Arrays;
import java.util.ArrayList;
import java.util.Collections;
import java.util.stream.Collectors;

class Pair {
    Integer first;
    Integer second;

    Pair(Integer first, Integer second) {
        this.first = first;
        this.second = second;
    }

    public String toString() {
        return "(" + first + ", " + second + ")";
    }
};

class Student {
    public static List<Pair> occurr(List<Integer> vet) {
         if(vet.isEmpty())
            return new ArrayList<>();
        
        List<Pair> list = new ArrayList<>();
        List<Integer> compare;
        vet.sort((a, b) -> Integer.compare(Math.abs(a), Math.abs(b)));
        compare = vet.stream().map(Math::abs).distinct().collect(Collectors.toList());
        
        for(Integer i : compare){
            Integer j = Collections.frequency(vet.stream().map(Math::abs).toList(), i);
            list.add(new Pair(i, j));
        }
        return list;
    }

    public static List<Pair> teams(List<Integer> vet) {
        if(vet.isEmpty())
            return new ArrayList<>();
        
        List<Pair> list = new ArrayList<>();
        List<Integer> compare;

        compare = vet.stream().map(Math::abs).collect(Collectors.toList());
        
        for(int i=0, count=1; i<compare.size(); i++){
            if(i<compare.size()-1 && compare.get(i).equals(compare.get(i+1))){
                count++;
            }
            else if (i==compare.size()-1){
                list.add(new Pair(compare.get(i), count));
                count = 1;
            }
            else{
                list.add(new Pair(compare.get(i), count));
                count = 1;
            }
        }
        return list;    
    }

    public static List<Integer> mnext(List<Integer> vet) {
         if(vet.isEmpty())
            return new ArrayList<>();
        
        List<Integer> list = new ArrayList<>();
        
        for(int i=0; i<vet.size(); i++){
            if(i<vet.size()-1 && vet.get(i) > 0 && vet.get(i+1) < 0){
                list.add(1);
            }
            else if(i>0 && vet.get(i) > 0 && vet.get(i-1) < 0){
                list.add(1);
            }
            else{
                list.add(0);
            }
        }
        return list;
    }

    private static boolean hasWoman(List<Integer> vet, int size, int index) {
        return true;
    }

    public static List<Integer> alone(List<Integer> vet) {
        if(vet.isEmpty())
            return new ArrayList<>();
        
        List<Integer> list = new ArrayList<>();
        
        for(int i=0; i<vet.size(); i++){
            if(vet.get(i)>0){
                if(vet.size()==1)
                    list.add(1);
                else if(i==0 && vet.get(i+1)>0)
                    list.add(1);
                else if(i==vet.size()-1 && vet.get(i-1)>0)
                    list.add(1);
                else if(i>0 && vet.get(i-1)>0 && vet.get(i+1)>0)
                    list.add(1);
                else
                    list.add(0);
            }
            else{
                list.add(0);
            }
        }
        return list;
    }

    public static int couple(List<Integer> vet) {
        return 0;
    }

    public static boolean hasSubseq(List<Integer> vet, List<Integer> seq, int pos) {
        return true;
    }

    public static int subseq(List<Integer> vet, List<Integer> seq) {
        return -1;
    }

    public static List<Integer> erase(List<Integer> vet, List<Integer> pos_list) {
        return new ArrayList<Integer>();
    }

    public static List<Integer> clear(List<Integer> vet, int value) {
        return new ArrayList<Integer>();
    }
}

public class Shell {
    public static void main(String[] _args) {
        while (true) {
            var line = scanner.nextLine();
            var args = line.split(" ");
            System.out.println('$' + line);

            if (args[0].equals("end")) {
                break;
            }
            else if (args[0].equals("occurr")) {
                var array = strToVet(args[1]);
                System.out.println("" + Student.occurr(array));
            }
            else if (args[0].equals("teams")) {
                var array = strToVet(args[1]);
                System.out.println("" + Student.teams(array));
            }
            else if (args[0].equals("mnext")) {
                var array = strToVet(args[1]);
                System.out.println("" + Student.mnext(array));
            }
            else if (args[0].equals("alone")) {
                var array = strToVet(args[1]);
                System.out.println("" + Student.alone(array));
            }
            else if (args[0].equals("erase")) {
                var array = strToVet(args[1]);
                var positions = strToVet(args[2]);
                System.out.println("" + Student.erase(array, positions));
            }
            else if (args[0].equals("clear")) {
                var array = strToVet(args[1]);
                var value = Integer.parseInt(args[2]);
                System.out.println("" + Student.clear(array, value));
            }
            else if (args[0].equals("subseq")) {
                var array = strToVet(args[1]);
                var seq = strToVet(args[2]);
                System.out.println("" + Student.subseq(array, seq));
            }
            else if (args[0].equals("couple")) {
                var array = strToVet(args[1]);
                System.out.println("" + Student.couple(array));
            }
            else {
                System.out.println("Invalid command");
            }
        }
    }

    public static List<Integer> strToVet(String s) {
        if (s.equals("[]"))
            return new ArrayList<>();
        List<String> parts = Arrays.asList(s.substring(1, s.length() - 1).split(","));
        List<Integer> result = new ArrayList<>();
        for (String part : parts)
            result.add(Integer.parseInt(part.trim()));
        return result;
    }

    static Scanner scanner = new Scanner(System.in);
}
