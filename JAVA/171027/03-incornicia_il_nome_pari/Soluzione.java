import java.util.*;
public class Soluzione {
	public static void main(String[] args) {
		Scanner in = new Scanner(System.in);
		String s = in.nextLine();
		int l = s.length();
		String st = "********************";
		String sp = "                    ";
		if (l % 2 != 0) {
			System.out.println("**" + st.substring(0, l) + "***");
			System.out.println("* " + sp.substring(0, l) + "  *");
			System.out.println("* " + s + "  *");
			System.out.println("* " + sp.substring(0, l) + "  *");
			System.out.println("**" + st.substring(0, l) + "***");
		} else {
			System.out.println("**" + st.substring(0, l) + "**");
			System.out.println("* " + sp.substring(0, l) + " *");
			System.out.println("* " + s + " *");
			System.out.println("* " + sp.substring(0, l) + " *");
			System.out.println("**" + st.substring(0, l) + "**");
		}
	}
}
