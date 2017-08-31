//Dependencies
import React, {Component} from 'react';

//Styles
import './RegisterUser.css';

//Material-UI
import TextField from 'material-ui/TextField';
import Paper from 'material-ui/Paper';
import RaisedButton from 'material-ui/RaisedButton';
import Done from 'material-ui/svg-icons/action/done';
//import axios from 'axios';

const styles = {

    paper: {
        width: '60%',
        height: '100%',
        marginLeft: '18.7%'

    },

    tittle: {
        paddingTop: 45,
        marginLeft: '10%',
        marginBottom: 12,
        fontWeight: 400,
    },

    fieldColor: {
        color: '#18BC9C',
    },

    labelColor: {
        color: '#000000',
    },

    buttonTextColor: {
        color: '#fff'
    },

    buttonCustumized: {
        width: 250,
    },
}

class RegisterUserUser extends Component {
    constructor() {
        super();
        this.state = {users: []}
    }

    /*componentWillMount(){
        axios.get('https://api.github.com/users')
            .then((response) => {
                //console.log(response);
                //localStorage.users = response.data;
                this.setState({
                    users: response.data,
                });
                //console.log(this.state.users)
            })
            .catch((error) => {
                console.log(error);
            })
    }*/

    /*handleFindUser(username) {
        return () => {
            axios.get('https://api.github.com/users/'+username)
                .then((response) => {
                    console.log(response.data)
                    alert("Hola soy "+response.data.login+" este es mi github")
                })
                .catch((error) => {
                    console.log(error);
                })
        }
    }*/

    render() {
        return (
            <div className={'RegisterUser-content center'}>
                <div className="RegisterUser-paper-form">
                    <Paper style={styles.paper} zDepth={5}>
                        {/*<div>
                            { this.state.users.map((value) => {
                                //let url_user = "https://api.github.com/users/"+value.login
                                return <div > <a href="#" onClick={this.handleFindUser(value.login)}>{value.login}</a> </div>
                            })
                            }
                        </div>*/}
                        <div className="RegisterUser-Form">
                            <h3 style={styles.tittle}>Register of user</h3>
                            <TextField
                                hintText={'User Name'}
                                required={true}
                                floatingLabelText={'User Name'}
                                type={'text'}
                                underlineFocusStyle={{borderColor: '#18BC9C'}}
                                floatingLabelFocusStyle={styles.fieldColor}
                                floatingLabelStyle={styles.labelColor}
                                className={'Input-One'}
                            />
                            <TextField
                                hintText={'Email'}
                                floatingLabelText={'Email'}
                                type={'text'}
                                underlineFocusStyle={{borderColor: '#18BC9C'}}
                                floatingLabelFocusStyle={styles.fieldColor}
                                floatingLabelStyle={styles.labelColor}
                                className={'Input-Two'}
                            />
                            <br />
                            <TextField
                                hintText={'Password'}
                                floatingLabelText={'Password'}
                                type={'password'}
                                underlineFocusStyle={{borderColor: '#18BC9C'}}
                                floatingLabelFocusStyle={styles.fieldColor}
                                floatingLabelStyle={styles.labelColor}
                                className={'Input-One'}
                            />
                            <TextField
                                hintText={'Confirm Password'}
                                floatingLabelText={'Confirm Password'}
                                type={'password'}
                                underlineFocusStyle={{borderColor: '#18BC9C'}}
                                floatingLabelFocusStyle={styles.fieldColor}
                                floatingLabelStyle={styles.labelColor}
                                className={'Input-Two'}
                            />
                            <br />
                        </div>

                        <div className={'RegisterUser-button'}>
                            <RaisedButton
                                label={'Register'}
                                backgroundColor={'#18BC9C'}
                                labelColor={'#fff'}
                                buttonStyle={styles.buttonTextColor}
                                floatingLabelStyle={styles.labelColor}
                                style={styles.buttonCustumized}
                                icon={<Done/>}
                            />
                        </div>

                    </Paper>
                </div>
            </div>

        );
    }
}
export default RegisterUserUser;