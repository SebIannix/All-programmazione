import java.util.*;
public class Soluzione {
	public static void main(String[] args) {
		Scanner in = new Scanner(System.in);
		String s = in.nextLine();
		int x = s.length();
		int d = s.charAt(x - 2) - '0';
		int u = s.charAt(x - 1) - '0';
		int sol = (2 * d) + u;
		if (sol == 4 || sol == 8 || sol == 12 || sol == 16 || sol == 20 || sol == 24) {
			System.out.println("divisibile");
		} else
			System.out.println("non divisibile");
	}
}
