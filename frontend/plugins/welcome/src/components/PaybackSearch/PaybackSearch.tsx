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
import { EntBank } from '../../api/models/EntBank';
import { EntAmountpaid } from '../../api/models/EntAmountpaid';
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
import { EntRecordinsurance } from '../../api/models/EntRecordinsurance';
import moment from 'moment';
import { EntPayback } from '../../api';

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


export default function Payback() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [paybacks, setPaybacks] = useState<EntPayback[]>([]);
  const [loading, setLoading] = useState(true);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = React.useState(true);
  const [alerttype, setAlertType] = useState(String);
  const [errormessege, setErrorMessege] = useState(String);
  const [officers, setOfficers] = useState<EntOfficer[]>([]);
  const [officerid, setOfficerID] = useState(Number);
  const [paybacksearch, setPaybackSearch] = useState(String);
 
  useEffect(() => {

    const getOfficers = async () => {
 
      const ofc = await api.listOfficer({});
        setLoading(false);
        setOfficers(ofc);
      };
      getOfficers();

      const dataa = localStorage.getItem("officerdata");
      if (dataa) {
      setOfficerID(Number(localStorage.getItem("officerdata")));
      setLoading(false);
      }

      const checkJobPosition = async () => {
        const jobdata = JSON.parse(String(localStorage.getItem("positiondata")));
        setLoading(false);
        if (jobdata != "พนักงานบริษัทประกันสุขภาพ" ) {
          localStorage.setItem("officerdata",JSON.stringify(null));
          localStorage.setItem("positiondata",JSON.stringify(null));
          history.pushState("","","./Officerlongin");  
          window.location.reload(false);       
        }
        else{
            setOfficerID(Number(localStorage.getItem("officerdata")))
        }
      }
    checkJobPosition();

  }, [loading]);

  const SearchPayback = async () => {
    const res = await api.listPayback({ offset: 0 });
    const search = Paybacksearch(res);
    setErrorMessege("ไม่พบเลขบัตรประชาชนที่ค้นหา");
        setAlertType("error");
        setAlert(false);
        setPaybacks([]);
        if(search.length > 0){
            Object.entries(check).map(([key, value]) =>{
                if (value == true){
                    setErrorMessege("ค้นหาข้อมูลสำเร็จ");
                    setAlertType("success");
                    setAlert(true);
                    setPaybacks(search);
                }
            })
        }

        setStatus(true);
  }

  const Paybacksearch = (res: any) => {
    const data = res.filter((filter: EntPayback) => filter?.paybackAccountiden?.includes(paybacksearch))
    if (data.length != 0 && paybacksearch != "") {
        return data;
    }
    else{
      return data;
        }
    }

  const handleSearchChange = (event: any) => {
    setPaybackSearch(event.target.value as string);
  };

    return (
      <Page theme={pageTheme.service}>
        <Header style={HeaderCustom} title={`ระบบค้นหาการคืนทุนประกัน`}>
        </Header>
        <Content>
        <ContentHeader title="ค้นหาการคืนทุนประกันด้วยเลขบัตรประชาชน">
        <Link component={RouterLink} to="/payback">
           <Button variant="contained" color="primary">
             บันทึกข้อมูลการคืนทุนประกัน
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
                  id="paybackContact"
                  label=""
                  variant="outlined"
                  color="secondary"
                  type="string"
                  size="small"
                  value={paybacksearch}
                  onChange={handleSearchChange}
                  name="payback_Contact"
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
              SearchPayback();
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
                                <TableCell align="center">id</TableCell>
                                <TableCell align="center">สมาชิกระบบประกันสุขภาพ</TableCell>
                                <TableCell align="center">เลขประจำตัวประชาชน</TableCell>
                                <TableCell align="center">ผลิตภัณฑ์</TableCell>
                                <TableCell align="center">ราคา</TableCell>
                                <TableCell align="center">ธนาคาร</TableCell>
                                <TableCell align="center">เลขบัญชี</TableCell>
                                <TableCell align="center">เบอร์โทร</TableCell>
                                <TableCell align="center">พนักงานที่ทำรายการ</TableCell>
                                <TableCell align="center">date</TableCell>
            </TableRow>
            </TableHead>
            <TableBody>
              {paybacks.map((item: any) => (
                <TableRow key={item.id}>
                <TableCell align="center">{item.id}</TableCell>
                <TableCell align="center">{item.edges?.member?.memberName}</TableCell>
                <TableCell align="center">{item.paybackAccountiden}</TableCell>
                <TableCell align="center">{item.edges?.product?.productName}</TableCell>
                <TableCell align="center">{item.edges?.product?.productPrice}</TableCell>
                <TableCell align="center">{item.edges?.bank?.bankType}</TableCell>
                <TableCell align="center">{item.paybackAccountnumber}</TableCell>
                <TableCell align="center">{item.paybackAccountname}</TableCell>
                <TableCell align="center">{item.edges?.officer?.officerName}</TableCell>
                <TableCell align="center">{moment(item.paybackTransfertime).format('DD/MM/YYYY HH:mm:ss')}</TableCell>
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