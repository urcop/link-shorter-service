import React, {useState} from 'react';
import Arrow from '../../assets/Arrow.svg'
import style from './Input.module.scss'
import axios from "axios";

const Input = () => {
    const [linkUrl, setLinkUrl] = useState('')
    async function createShort(event) {
        event.preventDefault()

        let newLink = {
            random: true,
            link: linkUrl
        }

        await axios
            .post('http://localhost:8080/api/v1/link', newLink)
            .then(response => console.log(response))
            .catch(err => console.log(err))
    }

    return (
        <div className={style.input_container}>
            <form onSubmit={createShort}>
                <input
                    type="text"
                    placeholder="Paste your link here..."
                    value={linkUrl}
                    onChange={(event) => setLinkUrl(event.target.value)}
                />
                <button type='submit'><img src={Arrow} alt="arrow"/></button>
            </form>
        </div>
    );
};

export default Input;