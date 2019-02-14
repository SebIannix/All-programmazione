import java.util.*;
public class soluzione {
	public static void main(String[] args) {
		Scanner in = new Scanner(System.in);
		
		String t = "";
		String s = t.toUpperCase();
		
		while (in.hasNextLine()) {
			s = in.nextLine().toUpperCase();
		
		String[] arr = s.split("\\W+");
		
		for (int i = 0; i < s.length(); i++) {
			for (int j = 0; j < alpha.length; j++) {
				if (s.charAt(i) == alpha[j])
					System.out.print(morse[j]);
			}
		}
	}

	}

	public static char[] alpha = { 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', ' ' };
	
	public static String[] morse = { ".- ", "-... ", "-.-. ", "-.. ", ". ", "..-. ", "--. ", ".... ", ".. ", ".--- ", "-.- ", ".-.. ", "-- ", "-. ", "--- ", ".--. ", "--.- ", ".-. ", "... ", "- ", "..- ", "...- ", ".-- ", "-..- ", "-.-- ", "--.. ", "" };
	
}
