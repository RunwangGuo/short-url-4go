import * as React from 'react';
import {DataGrid, type GridColDef, type GridRowSelectionModel} from '@mui/x-data-grid';
import Link from '@mui/material/Link';
import {useSafeState} from 'ahooks';
import {useEffect} from 'react';
import useService from '../service/service';
import {Button, Chip, TextField, Tooltip, Typography} from '@mui/material';
import {AccessAlarm, Add, ErrorOutlineOutlined} from '@mui/icons-material';
import AddFormDialog from './AddFormDialog';
import {
	DT_FORMAT,
	LINK_STATUS,
	LINK_STATUS_BUTTON_COLOR,
	LINK_STATUS_BUTTON_TEXT,
	LINK_STATUS_COLOR,
	LINK_STATUS_TEXT,
} from '../config/constants';
import {baseUrl} from '../api/api';
import moment from 'moment';
import ChangeStatusDialog from './ChangeStatusDialog';
import ChangeExpiredDialog from './ChangeExpiredDialog';
import SetRemarkDialog from './SetRemarkDialog';
import numeral from 'numeral';

export default function LinkTable() {
	const [links, setLinks] = useSafeState<API.LinkRowProp[]>([]);

	const [keyword, setKeyword] = useSafeState<string>();
	const [rowSelectionModel, setRowSelectionModel] = useSafeState<GridRowSelectionModel>([]);

	const [page, setPage] = useSafeState<number>(1);
	const [pageSize] = useSafeState<number>(30);
	const [total, setTotal] = useSafeState<number>(0);

	const [addVisible, setAddVisible] = useSafeState<boolean>(false);

	const [statusDialogVisible, setStatusDialogVisible] = useSafeState<boolean>(false);
	const [statusTargets, setStatusTargets] = useSafeState<string[]>([]);
	const [statusType, setStatusType] = useSafeState<number | undefined>();

	const [expiredDialogVisible, setExpiredDialogVisible] = useSafeState<boolean>(false);
	const [expiredTargets, setExpiredTargets] = useSafeState<string[]>([]);
	const [expiredValue, setExpiredValue] = useSafeState<number>();

	const [remarkDialogVisible, setRemarkDialogVisible] = useSafeState<boolean>(false);
	const [remarkTargets, setRemarkTargets] = useSafeState<string[]>([]);
	const [remarkValue, setRemarkValue] = useSafeState<string>();

	const {loading, search} = useService();

	const onSearch = (page?: number, keyword?: string) => {
		search({
			page: page ?? 1,
			size: pageSize,
			keyword: keyword?.trim(),
		}).then((res) => {
			setLinks(res.records ?? []);
			setTotal((res.size ?? 0) * (res.pages ?? 1));
		});
	};

	useEffect(() => {
		onSearch(1);
	}, []);

	const columns: Array<GridColDef<API.LinkRowProp>> = [
		{
			field: 'shortId',
			headerName: '短链接',
			width: 180,
			sortable: false,
			renderCell: (props) => {
				const url = `${baseUrl}/${props.row.shortId}`;
				return (
					<Typography
						variant={'body2'}
						color={'text.secondary'}
						align={'center'}
						display={'flex'}
						justifyContent={'space-between'}
					>
						<Link color={'inherit'} underline={'none'} href={url} target={'_blank'}>
							{url}
						</Link>
					</Typography>
				);
			},
		},
		{
			field: 'originalUrl',
			headerName: '原链接',
			minWidth: 320,
			cellClassName: 'cell-cls-name',
			sortable: false,
			renderCell: (props) => {
				return (
					<Typography
						variant={'body1'}
						variantMapping={{body1: 'div'}}
						color={'text.secondary'}
						display={'flex'}
						justifyContent={'space-between'}
						width={'100%'}
					>
						<Tooltip title={props.row.originalUrl} arrow={true} placement={'top'}>
							<Typography
								variant={'body2'}
								noWrap={true}
								textOverflow={'ellipsis'}
								width={'calc(100% - 20px)'}
							>
								<Link
									color={'inherit'}
									href={props.row.originalUrl}
									target={'_blank'}
									underline={'none'}
								>
									{props.row.originalUrl}
								</Link>
							</Typography>
						</Tooltip>
						{props.row.remark ? (
							<Tooltip title={props.row.remark} arrow={true} placement={'top'}>
								<ErrorOutlineOutlined fontSize={'small'} style={{cursor: 'pointer'}} />
							</Tooltip>
						) : null}
					</Typography>
				);
			},
		},
		{
			field: 'status',
			headerName: '状态',
			minWidth: 100,
			cellClassName: 'cell-cls-name',
			sortable: false,
			renderCell: (props) => {
				return (
					<Chip
						label={LINK_STATUS_TEXT[props.row.status ?? LINK_STATUS.UNKNOWN]}
						color={LINK_STATUS_COLOR[props.row.status ?? LINK_STATUS.UNKNOWN]}
						size={'small'}
					/>
				);
			},
		},
		{
			field: 'expiredTs',
			headerName: '有效期',
			minWidth: 160,
			cellClassName: 'cell-cls-name',
			sortable: false,
			valueGetter: ({value}) => (value ? moment(value).format(DT_FORMAT.DATETIME) : '永久'),
		},
		{
			field: 'hits',
			headerName: 'Hits',
			minWidth: 40,
			cellClassName: 'cell-cls-name',
			sortable: false,
			valueGetter: ({value}) => numeral(value ?? 0).format('0,0'),
		},
		{
			field: 'actions',
			type: 'actions',
			width: 240,
			renderCell: (props) => {
				return (
					<div className={'btn-actions-wrapper'}>
						<Button
							variant={'outlined'}
							size={'small'}
							color={LINK_STATUS_BUTTON_COLOR[props.row.status ?? LINK_STATUS.UNKNOWN]}
							onClick={() => {
								setStatusDialogVisible(true);
								setStatusType(
									props.row.status === LINK_STATUS.NORMAL
										? LINK_STATUS.DISABLED
										: LINK_STATUS.NORMAL,
								);
								setStatusTargets([props.row.shortId!]);
							}}
						>
							{props.row.status === LINK_STATUS.NORMAL
								? LINK_STATUS_BUTTON_TEXT[LINK_STATUS.DISABLED]
								: LINK_STATUS_BUTTON_TEXT[LINK_STATUS.NORMAL]}
						</Button>
						<Button
							variant={'outlined'}
							size={'small'}
							sx={{ml: 1}}
							onClick={() => {
								setExpiredDialogVisible(true);
								setExpiredTargets([props.row.shortId!]);
								setExpiredValue(props.row.expiredTs);
							}}
						>
              设置有效期
						</Button>
						<Button
							variant={'outlined'}
							size={'small'}
							sx={{ml: 1}}
							color={'secondary'}
							onClick={() => {
								setRemarkDialogVisible(true);
								setRemarkTargets([props.row.shortId!]);
								setRemarkValue(props.row.remark);
							}}
						>
              备注
						</Button>
					</div>
				);
			},
		},
	];

	return (
		<div className={'table-container'}>
			<div className={'table-body'}>
				<div className={'search-container'}>
					<TextField
						label={'请输入关键词'}
						variant={'standard'}
						inputMode={'text'}
						onChange={(event) => {
							setKeyword(event.target.value);
						}}
					/>
					<Button
						variant={'text'}
						size={'small'}
						onClick={() => {
							onSearch(1, keyword);
						}}
					>
            查询
					</Button>
				</div>
				<div className={'actions-container'}>
					{rowSelectionModel.length ? (
						<Button
							variant={'outlined'}
							startIcon={<AccessAlarm />}
							sx={{mr: 1}}
							onClick={() => {
								setExpiredDialogVisible(true);
								setExpiredTargets(rowSelectionModel.map((i) => i.toString()));
								setExpiredValue(undefined);
							}}
						>
              设置有效期
						</Button>
					) : null}
					<Button
						variant={'outlined'}
						startIcon={<Add />}
						color={'warning'}
						onClick={() => {
							setAddVisible(true);
						}}
					>
            添加
					</Button>
				</div>
			</div>

			<div style={{height: 600, width: '100%'}}>
				<DataGrid
					getRowId={(row) => row.shortId!}
					loading={loading}
					rows={links}
					columns={columns}
					rowCount={total}
					rowHeight={42}
					checkboxSelection={true}
					disableRowSelectionOnClick={true}
					initialState={{
						pagination: {
							paginationModel: {page: 0, pageSize},
						},
					}}
					rowSelectionModel={rowSelectionModel}
					onRowSelectionModelChange={(rowSelectionModel) => {
						setRowSelectionModel(rowSelectionModel);
					}}
					pageSizeOptions={[pageSize]}
					paginationMode={'server'}
					onPaginationModelChange={(model) => {
						setPage(model.page + 1);
						onSearch(model.page + 1, keyword);
					}}
				/>
			</div>
			<AddFormDialog
				visible={addVisible}
				onCancel={() => {
					setAddVisible(false);
				}}
				onOk={() => {
					setAddVisible(false);
					onSearch(1, keyword);
				}}
			/>
			<ChangeStatusDialog
				visible={statusDialogVisible}
				targets={statusTargets}
				status={statusType}
				onOk={() => {
					onSearch(page);
				}}
				onCancel={() => {
					setStatusDialogVisible(false);
					setStatusTargets([]);
					setStatusType(undefined);
				}}
			/>
			<ChangeExpiredDialog
				visible={expiredDialogVisible}
				targets={expiredTargets}
				defaultValue={expiredValue}
				onOk={() => {
					onSearch(page);
				}}
				onCancel={() => {
					setExpiredTargets([]);
					setExpiredDialogVisible(false);
				}}
			/>
			<SetRemarkDialog
				visible={remarkDialogVisible}
				targets={remarkTargets}
				defaultValue={remarkValue}
				onOk={() => {
					onSearch(page);
				}}
				onCancel={() => {
					setRemarkTargets([]);
					setRemarkDialogVisible(false);
				}}
			/>
		</div>
	);
}
