import React, { FC, useEffect,useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SearchIcon from '@material-ui/icons/Search';
import Swal from 'sweetalert2'; // alert

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Avatar,
  Button,
  FormHelperText,
} from '@material-ui/core'
import { DefaultApi } from '../../api/apis';
import Autocomplete from '@material-ui/lab/Autocomplete';
import { EntProduct } from '../../api/models/EntProduct';
import { EntMember } from '../../api/models/EntMember';
import { EntHospital } from '../../api/models/EntHospital';
import { EntOfficer } from '../../api/models/EntOfficer';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import InputAdornment from '@material-ui/core/InputAdornment';
import { Link as RouterLink } from 'react-router-dom';
import OutlinedInput from '@material-ui/core/OutlinedInput';
import { Alert } from '@material-ui/lab';
import { ContentHeader } from '@backstage/core';
import { Link } from '@material-ui/core';
import { Theme, createStyles } from '@material-ui/core/styles';

import CardMedia from '@material-ui/core/CardMedia';
import { Image1Base64Function } from '../../image/image1';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { EntInsurance } from '../../api/models/EntInsurance';
import moment from 'moment';

// header css
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
    width: 300,
  },
  table: {
    minWidth: 650,
  },
  margin: {
    margin: theme.spacing(1),
  },
  media: {
    height: 0,
    marginLeft: 25,
    maxWidth: 300,
}
}));

const check = {
  membercheck : true
}

const searchcheck = {
  resultsearchcheck: true
}

export default function Insurance() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [insurances, setInsurances] = useState<EntInsurance[]>([]);
  const [loading, setLoading] = useState(true);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = React.useState(true);
  const [alerttype, setAlertType] = useState(String);
  const [errormessege, setErrorMessege] = useState(String);
  const [members, setMembers] = useState<EntMember[]>([]);
  const [memberid, setMemberID] = useState(Number);
  const [insurancesearch, setInsuranceSearch] = useState(String);
 
  useEffect(() => {

    const getMembers = async () => {
 
      const m = await api.listMember({});
        setLoading(false);
        setMembers(m);
      };
      getMembers();

      const checkJobPosition = async () => {
        const jobdata = JSON.parse(String(localStorage.getItem("positiondata")));
        setLoading(false);
        if (jobdata != "สมาชิกระบบประกันสุขภาพ" ) {
          localStorage.setItem("memberdata",JSON.stringify(null));
          localStorage.setItem("positiondata",JSON.stringify(null));
          history.pushState("","","./SignIn");
          window.location.reload(false);        
        }
        else{
            setMemberID(Number(localStorage.getItem("memberdata")))
        }
      }
    checkJobPosition();

  }, [loading]);

  const SearchInsurance = async () => {
    const res = await api.listInsurance({ offset: 0 });
    const search = Insurancesearch(res);
    setErrorMessege("ไม่พบเลขประจำตัวประชาชนที่ค้นหา");
        setAlertType("error");
        setAlert(false);
        setInsurances([]);
        if(search.length > 0){
            Object.entries(check).map(([key, value]) =>{
                if (value == true){
                    setErrorMessege("พบข้อมูลการซื้อประกัน");
                    setAlertType("success");
                    setAlert(true);
                    setInsurances(search);
                }
            })
        }

        setStatus(true);
  }

  const Insurancesearch = (res: any) => {
    const data = res.filter((filter: EntInsurance) => filter?.insuranceIdentification?.includes(insurancesearch))
    if (data.length != 0 && insurancesearch != "") {
        return data;
    }
    else{
      return data;
        }
    }

  const handleSearchChange = (event: any) => {
    setInsuranceSearch(event.target.value as string);
  };

    return (
      <Page theme={pageTheme.tool}>
        <Header style={HeaderCustom} title={`ระบบการค้นหาประวัติการซื้อประกันสุขภาพ`}>
        </Header>
        <Content>
        <ContentHeader title="ค้นหาด้วยเลขประชาชน">
        <Link component={RouterLink} to="/insurance">
           <Button variant="contained" color="secondary">
             ซื้อประกันสุขภาพ
           </Button>
         </Link>
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success" onClose={() => { setStatus(false) }}>
                  {errormessege}
                </Alert>
              ) : (
                  <Alert severity="error" onClose={() => { setStatus(false) }}>
                    {errormessege}
                  </Alert>
                )}
            </div>
          ) : null}
     </ContentHeader>
      <Container maxWidth="xl">
        <Grid container spacing={3}>
          <Grid item xs={12}></Grid>
          <Grid item xs={3}>
          <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  //style={{ width: 300 }}
                  id="insuranceIdentification"
                  label=""
                  variant="outlined"
                  color="secondary"
                  type="string"
                  size="small"
                  value={insurancesearch}
                  onChange={handleSearchChange}
                  name="insurance_identification"
                  multiline
                  rows={1}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </FormControl>
          </Grid>
          <Grid item xs={4}>
          <Button
            onClick={() => {
              SearchInsurance();
            }}
            variant="contained"
            //color="primary"
            size="large"
            startIcon={<SearchIcon />}
          >
            ค้นหา
          </Button>
         </Grid>
        </Grid>
      </Container>
        
        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center">ลำดับประกันที่ซื้อ</TableCell>
                <TableCell align="center">ชื่อ Prodcut</TableCell>
                <TableCell align="center">ชื่อผู้รับผลประโยชน์</TableCell>
                <TableCell align="center">ราคา</TableCell>
                <TableCell align="center">โรงพยาบาล</TableCell>
                <TableCell align="center">วันเวลาที่ซื้อ</TableCell>
                <TableCell align="center">วันที่ต้องการจ่ายงวดแรก</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {insurances.map((item: any) => (
                <TableRow key={item.id}>
                    <TableCell align="center">{item.id}</TableCell>
                    <TableCell align="center">{item.edges?.product?.productName}</TableCell>
                    <TableCell align="center">{item.insuranceInsurer}</TableCell>
                    <TableCell align="center">{item.edges?.product?.productPrice}</TableCell>
                    <TableCell align="center">{item.edges?.hospital?.hospitalName}</TableCell>
                    <TableCell align="center">{moment(item.insuranceTimeBuy).format('DD/MM/YYYY HH.mm น.')}</TableCell>
                    <TableCell align="center">{moment(item.insuranceTimeFirstpay).format('DD/MM/YYYY')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
        <br></br>
        </Content>
      </Page>
    );
  }