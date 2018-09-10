import java.net.Socket;
import java.net.ServerSocket;
import java.io.BufferedReader;
import java.io.BufferedOutputStream;
import java.io.IOException;
import java.io.PrintWriter;
import java.io.InputStreamReader;
import java.util.StringTokenizer;


public class InvoiceApiServer implements Runnable {

	static final int PORT = 8080;

	private Socket connect;

	public InvoiceApiServer(Socket c) {
		connect = c;
	}

	public static void main (String[] args) {
		

		try {
			// Create a new connection
			ServerSocket serverConnect = new ServerSocket(PORT);
			System.out.println("Server started. Listening for connections on port: " + PORT + "\n");

			//	Start a server and 
			while (true) {
				InvoiceApiServer myServer = new InvoiceApiServer(serverConnect.accept());
				Thread thread = new Thread(myServer);
				thread.start();
			}
		} catch (IOException e) {
			System.err.println("Server Connection error : " + e.getMessage());
		}
	}

	@Override
	public void run() {
		BufferedReader in = null; 
		PrintWriter out = null; 
		BufferedOutputStream dataOut = null;
		String fileRequested = null;
		try {
			in = new BufferedReader(new InputStreamReader(connect.getInputStream()));
			out = new PrintWriter(connect.getOutputStream());
			dataOut = new BufferedOutputStream(connect.getOutputStream());

			String input = in.readLine();
			StringTokenizer parse = new StringTokenizer(input);
			String method = parse.nextToken().toUpperCase();

			String route = parse.nextToken().toLowerCase();

			// We support both GET & POST methods
			System.out.println("Request received with method " + method + " and route " + route + "\n");
		} catch (IOException e) {
			System.err.println("Server Connection error : " + e.getMessage());
		}
	}

}


