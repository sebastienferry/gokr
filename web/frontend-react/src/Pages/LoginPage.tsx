import AuthService from "../Services/AuthService";
import React, { useState } from 'react';

export default function LoginPage() {

    const [username, setUserName] = useState("");
    const [password, setPassword] = useState("");
    
    async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();
        const token = await AuthService.login(
            username,
            password
        );

        if (token) {
            window.location.href = '/';
        }
    }
    
    return (
        <>
            <section className="hero">
                <div className="hero-body">
                    <p className="title">                        
                    </p>
                    <p className="subtitle has-text-centered">
                        Glad to see you again! Please sign in.
                    </p>
                    <div className="content">
                        <div className="columns">
                            <div className="column"></div>
                            <div className="column is-half">
                                <form className="box" onSubmit={handleSubmit}>
                                    <div className="field">
                                        <label className="label">Email</label>
                                        <div className="control">
                                            <input className="input" type="email" placeholder="e.g. alex@example.com" onChange={e => setUserName(e.target.value)} />
                                        </div>
                                    </div>

                                    <div className="field">
                                        <label className="label">Password</label>
                                        <div className="control">
                                            <input className="input" type="password" placeholder="********" onChange={e => setPassword(e.target.value)} />
                                        </div>
                                    </div>

                                    <button className="button is-primary">Sign in</button>
                                </form>
                            </div>
                            <div className="column"></div>
                        </div>

                    </div>
                </div>
            </section>
        </>
    );
}