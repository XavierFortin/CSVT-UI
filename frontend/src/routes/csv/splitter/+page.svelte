<script lang="ts">
	import {
		Fileupload,
		Input,
		InputAddon,
		Label,
		ButtonGroup,
		Select,
		Toggle,
		Checkbox,
		Button
	} from 'flowbite-svelte';
	import { SplitFile, SystemDirectorySelect, SystemFileSelect } from '$lib/wailsjs/go/main/App';
	import { os } from '$lib/wailsjs/go/models';
	import { preventDefault } from 'svelte/legacy';
	let selectedFiles = $state<FileList | null>(null);
	let splitBy = $state<'lines' | 'files'>('lines');
	let splitByAmount = $state<number>(1);
	let delimitor = $state<'comma' | 'semicolon' | 'tab' | 'pipe'>('comma');
	let cleanColumns = $state<boolean>(true);
	$effect(() => {
		console.log('Selected files:', selectedFiles);
	});

	function openFileDialog() {
		SystemFileSelect()
			.then((result) => {
				console.log('Selected directory:', result);
			})
			.catch((error) => {
				console.error('Error selecting files:', error);
			});
	}

	function handleSubmit() {
		if (selectedFiles && selectedFiles.length > 0) {
			const file = selectedFiles[0];
			SplitFile(os.File.createFrom(file), splitByAmount, delimitor, cleanColumns)
				.then((result) => {
					console.log('File split successfully:', result);
					// Handle success (e.g., show a notification or update the UI)
				})
				.catch((error) => {
					console.error('Error splitting file:', error);
					// Handle error (e.g., show an error message)
				});
		} else {
			console.warn('No file selected');
			// Handle case where no file is selected
		}
	}
</script>

<Label for="csv_file" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white">
	Select CSV file
</Label>
<Button onclick={openFileDialog}>Choose File</Button>
<Fileupload
	type="file"
	accept=".csv"
	clearable
	id="csv_file"
	class="mb-2"
	bind:files={selectedFiles}
/>

<div class="mt-6">
	<div class="flex align-middle">
		<div>
			<Label for="split_by">Split By</Label>
			<ButtonGroup>
				<Input type="number" id="split_by" required bind:value={splitByAmount} />
				<InputAddon class="m-0 p-0">
					<Select bind:value={splitBy} class="w-30">
						<option value="lines">Lines</option>
						<option value="files">Files</option>
					</Select>
				</InputAddon>
				<InputAddon class="m-0 p-0">
					<Checkbox id="clean_columns_check" bind:checked={cleanColumns} />
					<Label for="clean_columns_check" class="w-30">Clean Columns ?</Label>
				</InputAddon>
			</ButtonGroup>
		</div>
	</div>
</div>

<div class="mt-6">
	<Label for="delimitor">Delimitor</Label>
	<Select bind:value={delimitor} id="delimitor" class="w-30">
		<option value="comma">Comma ( , )</option>
		<option value="semicolon">Semicolon ( ; )</option>
		<option value="tab">Tab</option>
		<option value="pipe">Pipe ( | )</option>
	</Select>
</div>

<Button onclick={handleSubmit}>Submit</Button>
