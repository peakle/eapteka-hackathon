import React, {Fragment} from 'react';
import PropTypes from 'prop-types';

import {
    Panel,
    PanelHeader,
    Header,
    Button,
    Group,
    Cell,
    Div,
    Avatar,
    IconButton,
    WriteBarIcon,
    Search, WriteBar, Title
} from '@vkontakte/vkui';
import {Icon28Search, Icon28SmileOutline, Icon28VoiceOutline} from "@vkontakte/icons";

const Home = ({id, go, fetchedUser}) => {

    const startRecording = () => {
        fetch('/v1/speech/state', {
            method: 'POST',
            headers: {
                'Content-Type': 'audio/ogg'
              },
            body: ''
        })
        .then((response) => console.log(response.url))
        .catch((e) => console.log(e));
    };

    return (
    <Panel id="writebar" centered>
        <Title level="1" weight="bold" style={{ padding: "20px 20px 60px", textAlign: "center" }}>
            Назовите лекарство
        </Title>
        <Group>
            {/* <WriteBar
                after={
                    <Fragment>
                        <WriteBarIcon onClick={() => startRecording()}>
                            <Icon28VoiceOutline />
                        </WriteBarIcon>
                        <WriteBarIcon>
                            <Icon28Search />
                        </WriteBarIcon>
                    </Fragment>
                }
                placeholder="Поиск"
            /> */}
        <IconButton onClick={() => startRecording()}>
            <Icon28VoiceOutline />
        </IconButton>
        </Group>
        <Button style={{ marginTop: "80px" }} mode="outline" size="l">
            Аптечка
        </Button>
    </Panel>
)
};

Home.propTypes = {
	id: PropTypes.string.isRequired,
	go: PropTypes.func.isRequired,
	fetchedUser: PropTypes.shape({
		photo_200: PropTypes.string,
		first_name: PropTypes.string,
		last_name: PropTypes.string,
		city: PropTypes.shape({
			title: PropTypes.string,
		}),
	}),
};

export default Home;
