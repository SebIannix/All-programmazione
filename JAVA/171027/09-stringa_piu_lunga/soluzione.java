import java.util.*;
public class soluzione {
	public static void main(String[]args) {
		Scanner in = new Scanner(System.in);
		
		String s = in.nextLine();
		String [] a = s.split(" ");
		int length = 0;
		String toPrint = null;
		
		for (String t : a) {
			if (t.length() > length ) {
				length = t.length();
				toPrint = t;
			}
		}
		
		System.out.println(toPrint);
	}
}
