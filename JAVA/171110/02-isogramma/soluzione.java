import java.util.*;
public class soluzione {

	public static String s;
	public static int iso = 0;

	public static void main(String[] args) {
		Scanner in = new Scanner(System.in);
		
		while (in.hasNextLine())
			s = in.nextLine();
		isogramma(s);
		
		System.out.println(iso);
	}
	
	public static void isogramma(String s) {
		
		String is = s.toLowerCase();
		
		String[] array = s.split("");
		Set<String> value = new HashSet<String>(Arrays.asList(array));
		
		if (s.length() == value.size())
			iso++;
		
	}
	
}
