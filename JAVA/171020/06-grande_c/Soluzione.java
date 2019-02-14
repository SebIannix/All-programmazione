import java.util.Scanner;
public class Soluzione {
	public static void main(String[] args) {
		Scanner in = new Scanner(System.in);
		int n = in.nextInt();
		String stars = "**********";
		String star = "*";
		String up = stars.substring(0,n);
		int i = n-2;
		System.out.println(up);
		for(int k=0;k<i;k++) {
		System.out.println(star);}
		System.out.print(up);
	}
}	
