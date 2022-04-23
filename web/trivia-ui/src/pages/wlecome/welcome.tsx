import * as React from 'react';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';

interface Props {
    onClick: () => void;
}

const Welcome: React.FC<Props> = (props: Props) => {
    return (
        <Box sx={{ width: '100%'}}>
            <Typography variant="h2" component="div" gutterBottom>
            🥳🥳🥳🥳🥳🥳🥳🥳🥳🥳 LA MULTI ANI MARIASIIIIIIIIII  🥳🥳🥳🥳🥳🥳🥳🥳🥳🥳
            </Typography>
            <Typography variant="h4" gutterBottom>
            Ca cadou (😛) de ziua ta, avand in vedere ca faci 30 de ani, anul acesta am decis ca ar fi bine sa primesti un cadou mai personalizat.
            </Typography>
            <Typography variant="h4" gutterBottom>
            Cum si tu ne resfeti mereu facand cadouri personalizate, am zis sa iti facem si tie unul anul asta (da, stiu, sunt mai misto ale tale, dar na, atat s-a putut 😑)
            </Typography>
            <Typography variant="h4" gutterBottom>
            Speram ca iti va placea acest joculet de trivia si ca te vei distra cu el 🤗
            </Typography>
            <Typography variant="h4" gutterBottom>
            Uite la ce urmeaza sa te bagi. O sa urmeze o serie de inntrebari foarte serioase si importante pentru politica la nivel global.
            </Typography>
            <Typography variant="h4" gutterBottom>
            Tu trebuie sa raspunzi corect la ele. Nu raspunzi corect, mai incerci.
            </Typography>
            <Typography variant="h4" gutterBottom>
            Daca raspunzi corect, o sa vezi 😛.
            </Typography>
            <Typography variant="h4" gutterBottom>
            Dupa ce ai raspuns corect, poti merge la intrebarea urmatoare!
            </Typography>
            <Typography variant="h4" gutterBottom>
            Acum, trebuie sa iei la cunostinta ca accepti termenele si conditiile care sunt pe alt undeva fara sa le citesti ca sa continui!
            </Typography>
            <Typography variant="h4" gutterBottom>
            Sport la joc!!
            </Typography>
            <div style={{
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center',
            }}>
                <Button
                variant="contained"
                size="large"
                color="secondary"
                onClick={props.onClick}
                >
                        Accept Termenele si Conditiile ca primaru
                </Button>
            </div>
        </Box>
    );
}

export default Welcome;