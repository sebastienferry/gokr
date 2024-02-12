export default function AboutPage() {
    return (
        <>            
            <section className="hero">
                <div className="hero-body">
                    <p className="title">
                        About
                    </p>
                    <p className="subtitle">
                        This is a simple example of a React app with a few pages.
                    </p>
                    <div className="content">
                        <p>
                            This sample uses the following:
                            <ul>
                                <li><a href="https://react.dev/">React</a></li>
                                <li><a href="https://reactrouter.com/en/main">React Router</a></li>
                                <li><a href="https://www.typescriptlang.org/">TypeScript</a></li>
                                <li><a href="https://bulma.io/">Bulma</a></li>
                            </ul>                        
                        </p>
                        <p>
                            It is a simple example of a web app with a few pages.
                            It has a menu with links to the home page, about page, login page, and signup page.
                        </p>
                    </div>                    
                </div>
            </section>
        </>
    );
}