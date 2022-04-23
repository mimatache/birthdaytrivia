import React, { useState } from 'react';
import './App.css';
import GameModal from './pages/game/game';
import TriviaService from './services/questionaire/questionaire';
import ImageService from './services/image/image';
import Welcome from './pages/wlecome/welcome';

const App = () =>{

  const [accepted, setAccepted] = useState(false)

  const acceptHandler = () => {
    setAccepted(true)
  }

  const trivia = new TriviaService("http://localhost:8080")
  const image = new ImageService("http://localhost:8080")
  if (!accepted) {
    return <Welcome onClick={acceptHandler} />
  }
  return <GameModal api={trivia} image={image}/>
}

export default App;
