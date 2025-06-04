<script lang="ts">
	import '../app.css';
	import {
		Navbar,
		NavBrand,
		NavHamburger,
		NavUl,
		NavLi,
		DropdownItem,
		Dropdown
	} from 'flowbite-svelte';
	import { ChevronDownOutline } from 'flowbite-svelte-icons';

	let { children } = $props();
</script>

<Navbar style="z-50" class="m-0 bg-white p-0 dark:bg-gray-600">
	<NavBrand href="/">
		<span class="self-center text-xl font-semibold whitespace-nowrap dark:text-white">Tools</span>
	</NavBrand>
	<NavHamburger />
	<NavUl>
		<NavLi class="cursor-pointer">
			CSV<ChevronDownOutline class="text-primary-800 ms-2 inline h-6 w-6 dark:text-white" />
		</NavLi>
		<Dropdown simple>
			<DropdownItem href="/csv/splitter">CSV Splitter</DropdownItem>
		</Dropdown>
	</NavUl>
</Navbar>

<!-- Main content -->
<div class="color-scheme-auto container mx-auto pt-5 dark:text-white">
	{@render children()}
</div>
