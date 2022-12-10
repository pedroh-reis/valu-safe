import React from 'react';
import axios from 'axios';
import moment from 'moment';
import PropTypes from 'prop-types';
import AppBar from '../components/molecules/appBar';
import TabsBar from '../components/organisms/tabsBar';
import { Colors, MAIN_TEXT_FONT } from '../config';
import SelectAction from '../components/organisms/selectAction';
import Login from '../components/organisms/login';
import Graph from '../components/organisms/graph';
import { GraphicEq } from '@mui/icons-material';
import GoBackTab from '../components/molecules/goBackTab';
import apiData from '../mock/mock-data.json';
import { CircularProgress, Grid } from '@mui/material';
import LockerControl from '../components/molecules/lockerControl';
var request = require('request');

class MainContainer extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      userData: {
        email: null,
      },
      loading: false,
      selectedScreen: 0, // 0 = locker 1 ; 1 = locker 2
      selectedSubscreen: 'selector', // 'selector', 'control' or 'monitor'
      graphData1: null,
      graphData2: null,
      unlocked1: null,
      unlocked2: null,
    };
  }

  componentDidMount = async () => {
    await this.updateStatusAndHistory();
  };

  handleSubscreenSelected = async (subscreen) => {
    this.setState((prevState) => ({
      ...prevState,
      selectedSubscreen: subscreen,
    }));
  };

  handleChangeLockerStatus = async (lockerId) => {
    // HERE GOES THE CODE TO HANDLE LOCKER STATUS CHANGE
    // CONNECT TO API HERE

    // UPDATE LOCKER 1 STATUS
    await axios({
      method: 'post',
      url: 'http://localhost:8080/locker',
      headers: {
        'Content-Type': 'application/json',
      },
      data: JSON.stringify({ id: String(lockerId) }),
    })
      .then((response) => 
        console.log(response.data)
      )
      .catch((err) => {
        alert('Erro ao carregar dados');
        console.log(err);
      });

      await this.updateStatusAndHistory();

    // if (lockerId === 0) {
    //   this.setState((prevState) => ({
    //     ...prevState,
    //     unlocked1: !prevState.unlocked1,
    //   }));
    // } else if (lockerId === 1) {
    //   this.setState((prevState) => ({
    //     ...prevState,
    //     unlocked2: !prevState.unlocked2,
    //   }));
    // }

    console.log(this.state);
  };

  updateStatusAndHistory = async () => {

    // UPDATE LOCKER 1 STATUS
    await axios({
      method: 'post',
      url: 'http://localhost:8080/locker/state',
      headers: {
        'Content-Type': 'application/json',
      },
      data: JSON.stringify({ id: '0' }),
    })
      .then((response) => 
        this.setState(prevState => ({
          ...prevState,
          unlocked1: response.data.isUnlocked,
        }))
      )
      .catch((err) => {
        alert('Erro ao carregar dados');
        console.log(err);
      });

    // UPDATE LOCKER 2 STATUS
    await axios({
      method: 'post',
      url: 'http://localhost:8080/locker/state',
      headers: {
        'Content-Type': 'application/json',
      },
      data: JSON.stringify({ id: '1' }),
    })
      .then((response) => 
        this.setState(prevState => ({
          ...prevState,
          unlocked2: response.data.isUnlocked,
        }))
      )
      .catch((err) => {
        alert('Erro ao carregar dados');
        console.log(err);
      });

      console.log(this.state);

    // UPDATE LOCKER 1 HISTORY

    await axios({
      method: 'post',
      url: 'http://localhost:8080/locker/stats',
      headers: {
        'Content-Type': 'application/json',
      },
      data: JSON.stringify({ id: '0', timeframe: 'day', value: 1 }),
    })
      .then((response) => 
        this.setState(prevState => ({
          ...prevState,
          graphData1: response.data.history.map((element) => {
            var hours = new Date(element.timestamp).getHours();
            var minutes = new Date(element.timestamp).getMinutes();
            hours = hours < 10 ? '0' + hours : hours;
            minutes = minutes < 10 ? '0' + minutes : minutes;
            var time = hours + ':' + minutes;
            return {
              status: element.unlocked ? 'Aberto' : 'Fechado',
              time: time,
            };
          }),
        }))
      )
      .catch((err) => {
        alert('Erro ao carregar dados');
        console.log(err);
      });

      // UPDATE LOCKER 2 HISTORY

    await axios({
      method: 'post',
      url: 'http://localhost:8080/locker/stats',
      headers: {
        'Content-Type': 'application/json',
      },
      data: JSON.stringify({ id: '1', timeframe: 'day', value: 1 }),
    })
      .then((response) => 
        this.setState(prevState => ({
          ...prevState,
          graphData2: response.data.history.map((element) => {
            var hours = new Date(element.timestamp).getHours();
            var minutes = new Date(element.timestamp).getMinutes();
            hours = hours < 10 ? '0' + hours : hours;
            minutes = minutes < 10 ? '0' + minutes : minutes;
            var time = hours + ':' + minutes;
            return {
              status: element.unlocked ? 'Aberto' : 'Fechado',
              time: time,
            };
          }),
        }))
      )
      .catch((err) => {
        alert('Erro ao carregar dados');
        console.log(err);
      });

      console.log(this.state);

    
  };

  handleScreenSelected = (event, newValue) => {
    this.setState((prevState) => ({
      ...prevState,
      selectedScreen: newValue,
    }));
  };

  handleLoginSubmit = (email, password) => {
    if (
      (email === 'bruno@valusafe.com' || email === 'pedro@valusafe.com' || email === 'carlos@valusafe.com') &&
      password === '123456'
    ) {
      this.setState((prevState) => ({
        ...prevState,
        userData: {
          email: email,
        },
      }));
    } else {
      alert('Email ou senha inválidos');
    }
  };

  // <<<<<<<< DATA METHODS >>>>>>>>

  // <<<<<<<<    RENDER METHOD    >>>>>>>

  render() {
    if (this.state.userData.email) {
      return (
        <React.Fragment>
          <AppBar handleLogOutButton={() => window.location.reload()} />
          <TabsBar handleTabSelected={this.handleScreenSelected} value={this.state.selectedScreen} />
          {this.state.loading && (
            <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '78vh' }}>
              <CircularProgress style={{ color: Colors.VALUSAFE_BLUE }} />
            </div>
          )}
          {!this.state.loading && this.state.selectedSubscreen === 'selector' && (
            <SelectAction
              handleControlSelected={async () => await this.handleSubscreenSelected('control')}
              handleMonitorSelected={async () => await this.handleSubscreenSelected('monitor')}
            />
          )}
          {!this.state.loading && this.state.selectedSubscreen === 'monitor' && this.state.selectedScreen === 0 && (
            <div>
              <Graph data={this.state.graphData1} lockerId={this.state.selectedScreen} />
              <GoBackTab handleGoBack={async () => await this.handleSubscreenSelected('selector')} />
            </div>
          )}
          {!this.state.loading && this.state.selectedSubscreen === 'monitor' && this.state.selectedScreen === 1 && (
            <div>
              <Graph data={this.state.graphData2} lockerId={this.state.selectedScreen} />
              <GoBackTab handleGoBack={async () => await this.handleSubscreenSelected('selector')} />
            </div>
          )}
          {!this.state.loading && this.state.selectedSubscreen === 'control' && this.state.selectedScreen === 0 && (
            <div>
              <LockerControl
                handleChangeLockerStatus={async () => await this.handleChangeLockerStatus(0)}
                lockerId={0}
                lockerUnlocked={this.state.unlocked1}
              />
              <GoBackTab handleGoBack={async () => await this.handleSubscreenSelected('selector')} />
            </div>
          )}
          {!this.state.loading && this.state.selectedSubscreen === 'control' && this.state.selectedScreen === 1 && (
            <div>
              <LockerControl
                handleChangeLockerStatus={async () => await this.handleChangeLockerStatus(1)}
                lockerId={1}
                lockerUnlocked={this.state.unlocked2}
              />
              <GoBackTab handleGoBack={async () => this.handleSubscreenSelected('selector')} />
            </div>
          )}
        </React.Fragment>
      );
    } else {
      return (
        <React.Fragment>
          <Login handleLoginSubmit={this.handleLoginSubmit} />
        </React.Fragment>
      );
    }
  }
}

export default MainContainer;

const {} = PropTypes;
MainContainer.propTypes = {};

// PRÓXIMOS PASSOS:
// 0. programar tela de login; DONE
// 0.1. enviar print para Pedro das 2 telas já criadas para a apresentação de terça; DONE
// 1. desenhar tela de controle;
// 2. programar tela de controle para ambas as portas;
// 3. adicionar gráfico à tela de monitoramento de ambas as portas com botão de voltar;
// 4. colocar lógica de login
// 5. colocar ação no botão de log out
// 6. enviar prints para o Pedro
// 7. limpar os arquivos que remetem ao cib
// 8. colocar ícone em miniatura na aba de navegação
// 9. subir no heroku e enviar link para o Pedro
// 10. conectar com back-end
