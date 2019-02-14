import java.util.Scanner;
public class Soluzione {
	public static void main(String[] args) {
		Scanner in = new Scanner(System.in);
		String s = in.nextLine();
		int length = s.length();
		int l4 = length+4;
		int l2 = length+2;
		String stars = "************************";
		String space = "                        ";
		String pre = "* ";
		String post = " *";
		String name = pre.concat(s).concat(post);
		String up = stars.substring(0,l4);
		String spazio = space.substring(0,length);
		String down = pre.concat(spazio).concat(post);
		System.out.println(up);
		System.out.println(down);
		System.out.println(name);
		System.out.println(down);
		System.out.println(up);
	}
}
