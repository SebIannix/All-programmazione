import java.util.*;
public class soluzione {
	public static void main(String[]args) {
		Scanner in = new Scanner(System.in);
		int n = 0;
		while (in.hasNext()){
			String p = in.next();
			int v = 0;
			int c = 0;
			
			for(int x = 0; x < p.length(); x++) {
				char q = p.charAt(x);
				if (q == 'a' || q == 'e' || q == 'i' || q == 'o' || q == 'u'){
					v++;
				}
				else {
					c++;
				}
			}
			if (v > c) {
				n++;
			}
		}
		System.out.println(n);
	}
}
