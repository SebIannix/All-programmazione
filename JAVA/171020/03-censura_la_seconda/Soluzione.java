import java.util.Scanner;
public class Soluzione {
	public static void main(String[] args) {
		Scanner in = new Scanner(System.in);
		String s = in.nextLine();
		String l = in.nextLine();
		String t = in.nextLine();
		String q = "####################";
		int lung = l.length();
		String seconda = q.substring(0,lung);
		System.out.print(s + " ");
		System.out.print(seconda + " ");
		System.out.print(t);
	}
}
