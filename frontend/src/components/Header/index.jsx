import React from 'react';
import style from './Header.module.scss'

const Header = () => {
    return (
        <div className={style.container}>
            <h1>
                <span>MAKE <br/>IT</span> SHORT<span>ER</span>
            </h1>
            <svg width="147" height="147" viewBox="0 0 147 147" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path fillRule="evenodd" clipRule="evenodd" d="M73.5 147C73.353 106.457 40.5025 73.6362 0 73.6362C40.5929 73.6362 73.5 40.6678 73.5 0C73.6463 40.543 106.497 73.3638 147 73.3638C106.406 73.3638 73.5 106.332 73.5 147Z" fill="#4F4F4F"/>
            </svg>
        </div>
    );
};

export default Header;