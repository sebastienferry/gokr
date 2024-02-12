import 'react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faHome, faBook } from '@fortawesome/free-solid-svg-icons'
import { faGithub } from '@fortawesome/free-brands-svg-icons'
import AuthService from '../Services/AuthService'

function AuthButtons() {
    
    async function handleLogout() {
        await AuthService.logout();
        window.location.href = '/';
    }
    
    if (AuthService.isAuthenticated()) {
        return (
            <>
                <div className="navbar-item">
                    <div className="is-vcentered">
                        <p className="is-size-7">Welcome, {AuthService.getUsername()}</p>
                    </div>
                </div>
                <div className="navbar-item">
                    <div className="buttons">
                        <a className="button is-light" onClick={handleLogout}>
                            Logout
                        </a>
                    </div>
                </div>
            </>
        )
    } else {
        return (<>
            <div className="navbar-item">
                <div className="buttons">
                    <a className="button is-light" href='/signup'>
                        Sign up
                    </a>
                    <a className="button is-primary" href='/login'>
                        <strong>Login</strong>
                    </a>
                </div>
            </div>
        </>)
    }
}

export default function Menu() {
    
    return (
        <nav className="navbar" role="navigation" aria-label="main navigation">
            <div className="navbar-brand">
                <a className="navbar-item has-text-weight-bold" href="#">
                    SAMPLE
                </a>

                <a role="button" className="navbar-burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
                    <span aria-hidden="true"></span>
                    <span aria-hidden="true"></span>
                    <span aria-hidden="true"></span>
                </a>
            </div>

            <div id="navbarBasicExample" className="navbar-menu">
                <div className="navbar-start">
                    <a className="navbar-item" href='/'>
                        <FontAwesomeIcon icon={faHome} />&nbsp;Home
                    </a>
                    <a className="navbar-item" href='/about'>
                        <FontAwesomeIcon icon={faBook} />&nbsp;About
                    </a>
                    <a className="navbar-item" href='#'>
                        <FontAwesomeIcon icon={faGithub} />&nbsp;Github
                    </a>
                </div>

                <div className="navbar-end">
                    <AuthButtons />
                </div>
            </div>
        </nav>
    )
}