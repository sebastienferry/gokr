/**
 * AuthService.tsx
 * This service provides methods for authenticating users.
 */

export default class AuthService {

    /**
     * Logs in a user.
     * @param username The username.
     * @param password The password.
     * @returns A promise that resolves to the user's token.
     */
    static async login(username: string | null, password: string | null): Promise<string> {
        // return fetch('http://localhost:8080/login', {
        //     method: 'POST',
        //     headers: {
        //         'Content-Type': 'application/json'
        //     },
        //     body: JSON.stringify({ username, password })
        // })
        // .then(data => data.json())
        console.log('AuthService.login', username, password);

        // In a real app, this would make a network request to a server.
        const token = {
            "username": username, "email": "joe@mail.com", "roles": [
                "read", "write", "delete"
            ]
        };
        
        const tokenString = JSON.stringify(token);
        console.log('AuthService.login', tokenString);
        this.setToken(tokenString);
        return Promise.resolve(tokenString);
    }

    /**
     * Logs out a user.
     */
    static async logout() {
        // In a real app, this would make a network request to a server.
        // We'll just return a resolved promise for now.        
        console.log('AuthService.logout');
        sessionStorage.removeItem('token');
    }

    /**
     * Signs up a user.
     * @param email The email.
     * @param username The username.
     * @param password The password.
     * @returns A promise that resolves to the user's token.
     */
    static async signup(email: string, username: string, password: string): Promise<string> {
        // In a real app, this would make a network request to a server.
        // We'll just return a fake token for now.
        console.log('AuthService.signup', email, username, password);
        return Promise.resolve('fake-token');
    }

    /**
     * Gets the current user's username.
     * @returns A promise that resolves to the user's username.
     */
    static getUsername(): string {
        // In a real app, this would get the username from local storage.
        // We'll just return a fake username for now.
        return 'fake-username'
    }

    /**
     * Gets the current user's email.
     * @returns A promise that resolves to the user's email.
     */
    static async getEmail(): Promise<string | null> {
        // In a real app, this would get the email from local storage.
        // We'll just return a fake email for now.
        return Promise.resolve('fake-email');
    }

    /**
     * Tests if the user is authenticated.
     * @returns A promise that resolves to true if the user is authenticated, and false otherwise.
     */
    static isAuthenticated(): boolean {
        // In a real app, this would check if the user's token is valid.
        // We'll just return true for now.        
        const token = sessionStorage.getItem("token");
        console.log('AuthService.isAuthenticated', token);
        return token !== null && token !== undefined && token !== "";         
    }

    /**
     * Sets the user's token.
     * @param userToken 
     */
    static setToken(userToken: string) {
        sessionStorage.setItem('token', JSON.stringify(userToken));
    }

    /**
     * Gets the user's token.
     * @returns The user's token.
     */
    static getToken(): string | null {
        const tokenString = sessionStorage.getItem('token');
        const userToken = JSON.parse(tokenString?.valueOf() || "");
        return userToken?.token
    }
}