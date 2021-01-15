import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import ApartmentIcon from '@material-ui/icons/Apartment';
import FaceIcon from '@material-ui/icons/Face';
import AppBar from '@material-ui/core/AppBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import HowToRegIcon from '@material-ui/icons/HowToReg';
import Typography from '@material-ui/core/Typography';
import { orange, red} from '@material-ui/core/colors';
import EnhancedEncryptionIcon from '@material-ui/icons/EnhancedEncryption';


import {
  Container,
  Grid,
  Table,
} from '@material-ui/core';
import { Link as RouterLink } from 'react-router-dom';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command

const HeaderCustom = {
  minHeight: '50px',
};
// css style 
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 1000,
  },
  secondary: {
    main: '#ff3d00',
  }
}));

  
      function a11yProps(index: any) {
        return {
          id: `scrollable-force-tab-${index}`,
          'aria-controls': `scrollable-force-tabpanel-${index}`,
        };
      }

      export default function ScrollableTabsButtonForce() {
        const classes = useStyles();
        const [value, setValue] = React.useState(0);
      
        const handleChange = (event: React.ChangeEvent<{}>, newValue: number) => {
          setValue(newValue);
        };

  return (

    <Page theme={pageTheme.app}>
      <Header style={HeaderCustom} title={`ระบบประกันสุขภาพ`}>
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={1}>
            <Grid item xs={8}>
   <Table >
     <tr><td></td>
       <td  align="center">
           <br></br>
         <EnhancedEncryptionIcon  style={{ color: red[500] , fontSize: 150}}  />
     <Typography component="h1" variant="h5">
     SE63 Health Insurance
     </Typography><br></br>
     </td>
     </tr>

     <tr><td></td><td>
         <hr ></hr></td>
     </tr>
     <Grid item xs={5}>
              <div className={classes.paper}> </div>
            </Grid>
     <tr><br></br></tr>
     <tr><td></td>
     <td>
     <AppBar position="static" color="default">
        <Tabs
          value={value}
          onChange={handleChange}
          variant="scrollable"
          scrollButtons="on"
          indicatorColor="secondary"
          textColor="secondary"
          aria-label="scrollable force tabs example"
        >
           <Tab label="Member Longin" icon={<FaceIcon style={{ color: red[500] }} />}  component={RouterLink} to="/signin"/>
           <Tab label="Officer Longin" icon={<HowToRegIcon style={{ color: red[500] }} />}   component={RouterLink} to="/officerlongin"/>
        </Tabs>
      </AppBar>
     </td>
     </tr>
     </Table>
     </Grid> 
            </Grid>
        </Container>
      </Content>
   
    </Page>
  );
  }
