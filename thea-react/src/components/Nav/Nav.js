//Dependencies
import React, { Component } from 'react';

//Styles
import './Nav.css'

//Assets
import Menu from '../Menu/Menu'
import Modal from '../Login/Login'

//Material-UI
import Appbar from 'material-ui/AppBar';
import FlatButton from 'material-ui/FlatButton';
import Person from 'material-ui/svg-icons/social/person';



class Nav extends Component{

	constructor(props) {
	  	super(props);
	  	this.state = {menu: false, modal: false};
	}

	handleToggleMenu = () => this.setState({menu: !this.state.menu});
	handleToggleModal = () => this.setState({modal: !this.state.modal});

	render(){
		return (
			<Appbar
				title={<span className={'title-nav'}>CVDI</span>}
				onLeftIconButtonTouchTap={this.handleToggleMenu}>

			    <Menu
			    	open={this.state.menu}
			    	onTouchTap={this.handleToggleMenu}
			    	onRequestChange={this.handleToggleMenu}
			    />

			    <Modal
			    	open={this.state.modal}
			    	onTouchTapCancel={this.handleToggleModal}
			    	onRequestClose={this.handleToggleModal}
			    />
			</Appbar>

		);
	}
}

export default Nav;
