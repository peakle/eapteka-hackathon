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
    Search, WriteBar
} from '@vkontakte/vkui';
import {Icon28Search, Icon28SmileOutline, Icon28VoiceOutline} from "@vkontakte/icons";

const Home = ({id, go, fetchedUser}) => (
    <Panel id="writebar">
        <Group>
            <WriteBar
                after={
                    <Fragment>
                        <WriteBarIcon>
                            <Icon28VoiceOutline />
                        </WriteBarIcon>
                        <WriteBarIcon>
                            <Icon28Search />
                        </WriteBarIcon>
                    </Fragment>
                }
                placeholder="Поиск"
            />
        </Group>
    </Panel>
);

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
