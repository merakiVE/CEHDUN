//Dependencies
import React from 'react'
import {Link} from 'react-router-dom';

//Styles
import './Menu.css'

//Material-UI
import Drawer from 'material-ui/Drawer'
import MenuItem from 'material-ui/MenuItem'
import SvgIcon from 'material-ui/SvgIcon';
import Divider from 'material-ui/Divider';
import Person from 'material-ui/svg-icons/social/person';
import Apps from 'material-ui/svg-icons/navigation/apps.js';
import {ListItem} from 'material-ui/List';

const HomeIcon = (props) => (
    <SvgIcon {...props}>
        <path d="M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z"/>
    </SvgIcon>
);

class Menu extends React.Component {

    render() {
        return (
            <div classID="nav">
                <Drawer
                    docked={false}
                    open={this.props.open}
                    onRequestChange={this.props.onRequestChange}>

                    <Link to={'/'}>
                        <MenuItem
                            primaryText={"MerakiVe"}
                            onTouchTap={this.props.onTouchTap}
                            leftIcon={<HomeIcon />}
                        />
                    </Link>

                    <Link to={'/projects'}>
                        <MenuItem
                            primaryText={'Generar Api'}
                            onTouchTap={this.props.onTouchTap}
                            leftIcon={<Apps />}
                        />
                    </Link>

                    <Divider />

                    <Link to={'/login'}>
                        <ListItem
                            primaryText={'Iniciar Sesion'}
                            onTouchTap={this.props.onTouchTap}
                            leftIcon={<Person />}
                        />
                    </Link>
                </Drawer>
            </div>
        );
    }
}

export default Menu;
