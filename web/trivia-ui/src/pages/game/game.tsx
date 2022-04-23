import React, {useState, useEffect} from 'react';
import Radio from '@mui/material/Radio';
import RadioGroup from '@mui/material/RadioGroup';
import FormControlLabel from '@mui/material/FormControlLabel';
import FormControl from '@mui/material/FormControl';
import FormHelperText from '@mui/material/FormHelperText';
import FormLabel from '@mui/material/FormLabel';
import Button from '@mui/material/Button';

import TriviaService, { QuestionResponse } from '../../services/questionaire/questionaire'
import ImageService from '../../services/image/image'



interface GameProps {
    api: TriviaService;
    image: ImageService;
}

const GameModal: React.FC<GameProps> = (props: GameProps) => {
    const [value, setValue] = useState('');
    const [error, setError] = useState(false);
    const [nextDisabled, setNextDisabled] = useState(true)
    const [helperText, setHelperText] = useState('Choose wisely');
    const [isAnswerCorrect, setIsAnswerCorrect] = useState<Boolean>(false)
    const [q, setQuestion] = useState<QuestionResponse>();
    const [img, setImg] = useState<string>();

    useEffect(() => {
        const fetchQuestion = () => {
            props.api.getQuestion()
                .then((response) => setQuestion(response))
                .catch((err) => {
                    setError(err.message);
                })
        }
        fetchQuestion()
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [])


    const handleRadioChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setValue((event.target as HTMLInputElement).value);
        setHelperText(' ');
        setError(false);
    };

    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const resp = props.api.submitReponse(+value)
        resp.then((response) => {
            if (response.isAnswerCorrect) {
                setHelperText('Nimerishi mai!');
                setError(false);
                props.image.getImage(response.image).then((imageBlob) => {
                    const imageObjectURL = URL.createObjectURL(imageBlob);
                    setImg(imageObjectURL);
                })
                setIsAnswerCorrect(true)
                setNextDisabled(!response.hasNext)
            } else if (!response.isAnswerCorrect) {
                handleReset();
                setHelperText('Nu prea te pricepi. O luam de la capat!');
                setError(true);
            } else {
                setHelperText('Bafta :) !');
                setError(true);
            }
        }).catch((err) => {
            setError(err.message);
        })
        
    };

    const handleNextQuestion = () => {
        props.api.getQuestion()
            .then((response) => setQuestion(response))
            .catch((err) => {
                setError(err.message);
            })
        setValue('')
        setNextDisabled(true)
        setIsAnswerCorrect(false)
        setHelperText('')
    };

    const handleReset = () => {
        props.api.reset()
        handleNextQuestion()
    };




    return (
        <div style={{
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            fontSize: 10,
        }}>
        <form onSubmit={handleSubmit}>
        <FormControl sx={{ m: 3 }} error={error} variant="standard">
            <FormLabel id="demo-error-radios">{q?.question}</FormLabel>
            <RadioGroup
                aria-labelledby="demo-error-radios"
                name="quiz"
                value={value}
                onChange={handleRadioChange}
            >
                {q?.answers.map((v, i) => {
                    return <FormControlLabel value={i} control={<Radio />} label={v} key={i} />
                })}
            </RadioGroup>
            <FormHelperText>{helperText}</FormHelperText>
            <Button sx={{ mt: 1, mr: 1 }} type="submit" variant="outlined" style={{marginBottom: 10}}>
                Check Answer
            </Button>
            <Button
                variant="contained"
                size="large"
                color="secondary"
                onClick={handleNextQuestion}
                disabled={nextDisabled}
                style={{marginBottom: 10}}
            >
                Next Question
            </Button>
            <Button
                variant="contained"
                size="large"
                color="secondary"
                onClick={handleReset}
                style={{marginBottom: 10}}
            >
                Reset
            </Button>
        </FormControl>
        </form>
        {(isAnswerCorrect) && <img src={img} alt="icons" />}
        </div>
    );
}

export default GameModal;