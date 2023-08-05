import './App.css'
import React, { useState } from 'react';
import urlShotenerService from './services/urlShortenerService';
import {ToastContainer, toast} from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';


function App() {

  const [url, setUrl] = useState<string>("");
  const [shortenedUrl, setShortenedUrl] = useState<string>("");

  const urlShortenerService = new urlShotenerService("http://localhost:8080");

  const onUrlChange = (event:React.ChangeEvent<HTMLElement>) => {
    let currentUrlContent: string | null = event.target.value;
    if(currentUrlContent)
    {
      setUrl(currentUrlContent);
    }
  }

  const onGenerateShortUrlClick = async (event:React.MouseEvent<HTMLElement>) => {
    let trimmedUrl: string = url.trim();
    if(!trimmedUrl)
    {
      toast('empty link.')
      return
    }

    let [urlResult, error] = await urlShortenerService.generateUrl(url);
    let errorToastMessage = "there was a unknown problem. try again later!";
    if(error && error.length > 0)
    {
      errorToastMessage = error;
      toast(errorToastMessage)
    }

    if(!urlResult) {
      return
    }

    setShortenedUrl(urlResult);
  }

  return (
    <React.Fragment>
      <div className="page-heading-container">
        <div className="page-heading">
          <h1>URL Shortener</h1>
          <div className="input-container">
            <input id="generate-url-input" placeholder='https://example.com' onChange={onUrlChange} type='url'/>
            <button id="generate-url-button" onClick={onGenerateShortUrlClick}>short me!</button>
          </div>
          {
            shortenedUrl.length > 0 &&
            <h2 className="generated-url"><a href={shortenedUrl}>{shortenedUrl}</a></h2>
          }
        </div>
        <ToastContainer
        position="top-right"
        autoClose={1500}
        hideProgressBar={true}
        newestOnTop={false}
        closeOnClick
        rtl={false}
        theme="light"
        />
      </div>
    </React.Fragment>
  )
}

export default App
