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
import { EntMember } from '../../api/models/EntMember';
import { EntCategory } from '../../api/models/EntCategory';
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
import { EntInquiry } from '../../api/models/EntInquiry';
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


export default function Inquiry() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [inquiries, setInquiries] = useState<EntInquiry[]>([]);
  const [loading, setLoading] = useState(true);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = React.useState(true);
  const [alerttype, setAlertType] = useState(String);
  const [errormessege, setErrorMessege] = useState(String);
  const [members, setMembers] = useState<EntMember[]>([]);
  const [memberid, setMemberID] = useState(Number);
  const [inquirysearch, setInquirySearch] = useState(String);
 
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



  const SearchInquiry = async () => {
    const res = await api.listInquiry({ offset: 0 });
    console.log(res);
    const search = Inquirysearch(res);
    setErrorMessege("ไม่พบชื่อผู้สอบถามการสอบถามข้อมูลที่ค้นหา");
        setAlertType("error");
        setAlert(false);
        setInquiries([]);
        if(search.length > 0){
            Object.entries(check).map(([key, value]) =>{
                if (value == true){
                    setErrorMessege("พบข้อมูลการสอบถาม");
                    setAlertType("success");
                    setAlert(true);
                    setInquiries(search);
                }
            })
        }

        setStatus(true);
  }
  const Inquirysearch = (res: any) => {
    const data = res.filter((filter: EntInquiry) => filter?.inquiryNameMessages?.includes(inquirysearch))
    if (data.length != 0 && inquirysearch != "") {
        return data;
    }
    else{
      return data;
        }
    }

  const handleSearchChange = (event: any) => {
    setInquirySearch(event.target.value as string);
  };

  return (
    <Page theme={pageTheme.app}>
      <Header style={HeaderCustom} title={`ระบบค้นหาประวัติการสอบถามข้อมูล`}>
      </Header>
      <Content>
      <ContentHeader title="ค้นหาด้วยชื่อผู้สอบถาม">
      <Link component={RouterLink} to="/inquiry">
         <Button variant="contained" color="secondary">
           ติดต่อสอบถามข้อมูล
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
                id="inquiryNameMessages"
                label=""
                variant="outlined"
                color="secondary"
                type="string"
                size="small"
                value={inquirysearch}
                onChange={handleSearchChange}
                name="inquiry_name_messages"
                multiline
              />
            </FormControl>
        </Grid>
        <Grid item xs={4}>
        <Button
          onClick={() => {
            SearchInquiry();
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
              <TableCell align="center">ลำดับการสอบถาม</TableCell>
              <TableCell align="center">ชื่อผู้สอบถาม</TableCell>
              <TableCell align="center">อายุ</TableCell>
              <TableCell align="center">เบอร์โทรติดต่อกลับ</TableCell>
              <TableCell align="center">รายการสอบถาม</TableCell>
              <TableCell align="center">ชื่อผลิตภัณฑ์</TableCell>
              <TableCell align="center">รายละเอียด</TableCell>
              <TableCell align="center">วันเวลาที่สอบถาม</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {inquiries.map((item: any) => (
              <TableRow key={item.id}>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.inquiryNameMessages}</TableCell>
                  <TableCell align="center">{item.inquiryAgeMessages}</TableCell>
                  <TableCell align="center">{item.inquiryPhoneMessages}</TableCell>
                  <TableCell align="center">{item.edges?.category?.categoryName}</TableCell>
                  <TableCell align="center">{item.edges?.product?.productName}</TableCell>
                  <TableCell align="center">{item.inquiryMessages}</TableCell>
                  <TableCell align="center">{moment(item.inquiryTimeMessages).format('DD/MM/YYYY HH.mm น.')}</TableCell>
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