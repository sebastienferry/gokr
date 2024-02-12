export default function SingupPage() {
    return (
        <>
            <section className="hero">
                <div className="hero-body">
                    <p className="title">
                    </p>
                    <p className="subtitle has-text-centered">
                        Happy to see you! Please sign up.
                    </p>
                    <div className="content">
                        <div className="columns">
                            <div className="column"></div>
                            <div className="column is-half">
                                <form className="box">
                                    <div className="field">
                                        <label className="label">Email</label>
                                        <div className="control">
                                            <input className="input" type="email" placeholder="e.g. alex@example.com" />
                                        </div>
                                    </div>

                                    <div className="field">
                                        <label className="label">Username</label>
                                        <div className="control">
                                            <input className="input" type="text" placeholder="alex" />
                                        </div>
                                    </div>

                                    <div className="field">
                                        <label className="label">Password</label>
                                        <div className="control">
                                            <input className="input" type="password" placeholder="********" />
                                        </div>
                                    </div>

                                    <div className="field">
                                        <label className="label">Confirm password</label>
                                        <div className="control">
                                            <input className="input" type="password" placeholder="********" />
                                        </div>
                                    </div>

                                    <div className="level">
                                        <button className="button is-primary">Sign up</button>
                                        <a className="button is-light" href='/login'>I already have an account</a>
                                    </div>
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